package controller

import (
	"encoding/json"
	"fmt"
	"gin-research-sys/internal/form"
	"gin-research-sys/internal/model"
	"gin-research-sys/internal/request"
	"gin-research-sys/internal/response"
	"gin-research-sys/internal/service"
	"gin-research-sys/internal/util"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type IResearchController interface {
	List(c *gin.Context)
	Create(c *gin.Context)
	Retrieve(c *gin.Context)
	RetrieveOpen(c *gin.Context)
	Update(c *gin.Context)
	Destroy(c *gin.Context)

	Square(c *gin.Context)

	DownloadExcel(ctx *gin.Context)
}
type ResearchController struct{}

func NewResearchController() IResearchController {
	return ResearchController{}
}

var researchServices = service.NewResearchService()
var openRecordServices = service.NewOpenRecordService()

func (r ResearchController) List(ctx *gin.Context) {
	pagination := form.Pagination{}
	if err := ctx.ShouldBindQuery(&pagination); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, nil, "参数错误")
		return
	}
	var researches []model.Research
	page, size := pagination.Page, pagination.Size
	var total int64

	// 不是超管只可查看自己发布的问卷
	query := make(map[string]interface{})
	claims := jwt.ExtractClaims(ctx)
	id := int(claims["id"].(float64))
	if username := claims["username"].(string); username != "admin" {
		query["publisher_id"] = id
	}
	if err := researchServices.List(&researches, page, size, &total, query); err != nil {
		log.Println(err.Error())
		util.Success(ctx, nil, "获取数据失败")
		return
	}
	util.Success(ctx, gin.H{
		"page":    pagination.Page,
		"size":    pagination.Size,
		"results": researches,
		"total":   total,
	}, "获取数据成功")
}

func (r ResearchController) Retrieve(ctx *gin.Context) {
	var id int
	var err error
	if id, err = strconv.Atoi(ctx.Param("id")); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "ID错误")
		return
	}
	research := model.Research{}
	if err = researchServices.Retrieve(&research, id); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "未找到记录")
		return
	}
	util.Success(ctx, gin.H{"research": research}, "获取信息成功")
}

func (r ResearchController) RetrieveOpen(ctx *gin.Context) {
	var id int
	var err error
	if id, err = strconv.Atoi(ctx.Param("id")); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "ID错误")
		return
	}
	research := model.Research{}
	if err = researchServices.Retrieve(&research, id); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "未找到记录")
		return
	}
	if research.Open == 0 {
		util.Fail(ctx, gin.H{}, "非开放问卷")
		return
	}
	util.Success(ctx, gin.H{"research": research}, "获取信息成功")
}

func (r ResearchController) Create(ctx *gin.Context) {
	createForm := form.ResearchCreateForm{}
	if err := ctx.ShouldBindJSON(&createForm); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "参数错误")
		return
	}

	claims := jwt.ExtractClaims(ctx)
	id := int(claims["id"].(float64))
	research := model.Research{
		Title:       createForm.Title,
		Description: createForm.Description,
		Config:      createForm.Config,
		Items:       createForm.Items,
		Values:      createForm.Values,
		StartAt:     createForm.StartAt,
		EndAt:       createForm.EndAt,
		Access:      createForm.Access,
		Once:        createForm.Once,
		Open:        createForm.Open,
		PublisherID: id,
	}

	if err := researchServices.Create(&research); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "创建失败")
		return
	}
	util.Success(ctx, gin.H{}, "创建成功")
}

func (r ResearchController) Update(ctx *gin.Context) {
	var id int
	var err error
	id, err = strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "ID错误")
		return
	}
	updatePayload := request.ResearchUpdatePayload{}
	if err = ctx.ShouldBindJSON(&updatePayload); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "参数错误")
		return
	}
	research := model.Research{}
	if err = researchServices.Retrieve(&research, id); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "未找到记录")
		return
	}
	payload := map[string]interface{}{
		"title":       updatePayload.Title,
		"description": updatePayload.Description,
		"config":      updatePayload.Config,
		"start_at":    updatePayload.StartAt,
		"end_at":      updatePayload.EndAt,
		"access":      updatePayload.Access,
		"once":        updatePayload.Once,
	}
	if err = researchServices.Update(&research, payload); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "更新失败")
		return
	}
	util.Success(ctx, gin.H{}, "更新成功")
}

func (r ResearchController) Destroy(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "参数错误")
		return
	}
	if err = researchServices.Destroy(id); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "删除失败")
		return
	}
	util.Success(ctx, gin.H{}, "删除成功")
}

func (r ResearchController) Square(ctx *gin.Context) {
	pagination := form.Pagination{}
	if err := ctx.ShouldBindQuery(&pagination); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, nil, "参数错误")
		return
	}

	claims := jwt.ExtractClaims(ctx)
	id := int(claims["id"].(float64))
	user := model.User{}
	if err := userServices.Retrieve(&user, id); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "未找到记录")
		return
	}

	var researches []response.ResearchResponse
	page, size := pagination.Page, pagination.Size
	var total int64

	query := make(map[string]interface{})
	query["access"] = user.College
	if err := researchServices.FindByAccess(&researches, page, size, &total, query); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, nil, "获取数据失败")
		return
	}
	util.Success(ctx, gin.H{
		"page":    pagination.Page,
		"size":    pagination.Size,
		"results": researches,
		"total":   total,
	}, "获取数据成功")
}

func (r ResearchController) DownloadExcel(ctx *gin.Context) {
	var id int
	var err error
	if id, err = strconv.Atoi(ctx.Param("id")); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "ID错误")
		return
	}

	research := model.Research{}
	if err = researchServices.Retrieve(&research, id); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "获取信息失败")
		return
	}

	var items []map[string]interface{}
	if err = json.Unmarshal([]byte(research.Items), &items); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "数据解析失败")
		return
	}
	titles := []interface{}{"序号"}
	var fieldIDs []string
	for _, item := range items {
		fmt.Println(item)
		titles = append(titles, item["label"].(string))
		fieldIDs = append(fieldIDs, item["fieldID"].(string))
	}
	titles = append(titles, "IP地址", "填写时间")

	var openRecords []model.OpenRecord
	if err = openRecordServices.ListByResearchID(&openRecords, research.ID); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, nil, "获取数据失败")
		return
	}
	// 1. start generate excel
	xlsx := excelize.NewFile()
	// 2. new StreamWriter
	streamWriter, _ := xlsx.NewStreamWriter("Sheet1")
	if _, err = xlsx.NewStyle(`{"font":{"color":"#777777"}}`); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, nil, "设置样式失败")
		return
	}
	// 3. write title
	if err = streamWriter.SetRow("A1", titles); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, nil, "写入Row失败")
		return
	}

	// 4. write record data
	for k, v := range openRecords {
		row := []interface{}{k + 1}
		values := make(map[string]interface{})
		if err = json.Unmarshal([]byte(v.Values), &values); err != nil {
			log.Println(err.Error())
			util.Fail(ctx, gin.H{}, "数据解析失败")
			return
		}
		for _, fieldID := range fieldIDs {
			row = append(row, values[fieldID])
		}
		row = append(row, v.IPAddress)
		row = append(row, v.CreatedAt.Format("2006-01-02 15:04:05"))

		cell, _ := excelize.CoordinatesToCellName(1, k+2)
		if err = streamWriter.SetRow(cell, row); err != nil {
			log.Println(err.Error())
			util.Fail(ctx, nil, "写入Row失败")
			return
		}
	}
	// 5. flush streamWriter
	if err = streamWriter.Flush(); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, nil, "刷新失败")
		return
	}
	ctx.Header("response-type", "blob")
	data, _ := xlsx.WriteToBuffer()
	ctx.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", data.Bytes())
}
