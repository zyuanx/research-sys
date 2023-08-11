package controller

import (
	"errors"
	"gin-research-sys/internal/util"

	"log"
	"strconv"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/zyuanx/research-sys/internal/form"
	"github.com/zyuanx/research-sys/internal/model"
	"github.com/zyuanx/research-sys/internal/request"
	"github.com/zyuanx/research-sys/internal/service"
	"gorm.io/gorm"
)

type IRecordController interface {
	List(ctx *gin.Context)
	Create(ctx *gin.Context)
	Retrieve(ctx *gin.Context)
	Filled(ctx *gin.Context)
}
type RecordController struct{}

func NewRecordController() IRecordController {
	return RecordController{}
}

var recordServices = service.NewRecordService()

func (r RecordController) List(ctx *gin.Context) {
	pagination := form.Pagination{}
	if err := ctx.ShouldBindQuery(&pagination); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, nil, "参数错误")
		return
	}

	var records []model.Record
	var total int64

	// not admin, retrieve created by yourself
	query := make(map[string]interface{})
	claims := jwt.ExtractClaims(ctx)
	id := int(claims["id"].(float64))
	username := claims["username"].(string)
	if username != "admin" {
		query["publisher"] = id
	}
	if err := recordServices.List(pagination.Page, pagination.Size, &records, &total, query); err != nil {
		util.Success(ctx, nil, err.Error())
		return
	}
	util.Success(ctx, gin.H{
		"page":    pagination.Page,
		"size":    pagination.Size,
		"results": records,
		"total":   total,
	}, "")
}

func (r RecordController) Retrieve(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "参数错误")
		return
	}
	record := model.Record{}
	if err = recordServices.Retrieve(&record, id); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "获取记录失败")
		return
	}

	util.Success(ctx, gin.H{"record": record}, "获取成功")
}

func (r RecordController) Create(ctx *gin.Context) {
	createPayload := request.RecordCreatePayload{}
	if err := ctx.ShouldBindJSON(&createPayload); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "参数错误")
		return
	}

	// 获取用户信息
	claims := jwt.ExtractClaims(ctx)
	userID := int(claims["id"].(float64))

	// 获取问卷信息
	research := model.Research{}
	if err := researchServices.Retrieve(&research, createPayload.ResearchID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			util.Fail(ctx, gin.H{}, "未找到记录")
			return
		}
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "获取数据失败")
		return
	}
	// 检查是否在时间范围内
	now := time.Now()
	if now.Before(research.StartAt) || now.After(research.EndAt) {
		util.Fail(ctx, gin.H{}, "不在时间范围内")
		return
	}

	record := model.Record{
		ResearchID: createPayload.ResearchID,
		Values:     createPayload.Values,
		IPAddress:  ctx.ClientIP(),
		UserID:     userID,
	}
	if err := recordServices.Create(&record, &research); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "创建失败")
		return
	}
	util.Success(ctx, gin.H{}, "创建成功")
}

func (r RecordController) Filled(ctx *gin.Context) {
	researchId := ctx.Param("id")

	claims := jwt.ExtractClaims(ctx)
	id := int(claims["id"].(float64))

	record := model.Record{}
	if err := recordServices.FindByResearchID(&record, researchId, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			util.Success(ctx, gin.H{}, "")
			return
		}
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "retrieve fail")
		return
	}
	util.Fail(ctx, gin.H{}, "already filled in")
}

type IOpenRecordController interface {
	List(ctx *gin.Context)
	Create(ctx *gin.Context)
	Retrieve(ctx *gin.Context)
}
type OpenRecordController struct{}

func NewOpenRecordController() IOpenRecordController {
	return OpenRecordController{}
}

var openRecordService = service.OpenRecordService{}

func (o OpenRecordController) List(ctx *gin.Context) {
	listQuery := request.OpenRecordListQuery{}
	if err := ctx.ShouldBindQuery(&listQuery); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, nil, "参数错误")
		return
	}

	var records []model.OpenRecord
	page, size := listQuery.Page, listQuery.Size
	var total int64
	query := make(map[string]interface{})
	query["research_id"] = listQuery.ResearchID
	if err := openRecordService.List(page, size, &records, &total, query); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, nil, "获取错误")
		return
	}
	util.Success(ctx, gin.H{
		"page":    page,
		"size":    size,
		"results": records,
		"total":   total,
	}, "")
}
func (o OpenRecordController) Create(ctx *gin.Context) {
	createPayload := request.OpenRecordCreatePayload{}
	var err error
	if err = ctx.ShouldBindJSON(&createPayload); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "参数错误")
		return
	}
	openRecord := model.OpenRecord{
		ResearchID: createPayload.ResearchID,
		IPAddress:  ctx.ClientIP(),
		Values:     createPayload.Values,
	}
	if err = openRecordService.Create(&openRecord); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "创建失败")
		return
	}
	util.Success(ctx, gin.H{}, "创建成功")

}
func (o OpenRecordController) Retrieve(ctx *gin.Context) {
	var id int
	var err error
	if id, err = strconv.Atoi(ctx.Param("id")); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "ID错误")
		return
	}
	openRecord := model.OpenRecord{}
	if err = openRecordService.Retrieve(&openRecord, id); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "未找到记录")
		return
	}
	util.Success(ctx, gin.H{"openRecord": openRecord}, "获取成功")
}
