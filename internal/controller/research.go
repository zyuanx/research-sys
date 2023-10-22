package controller

import (
	"encoding/json"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zyuanx/research-sys/internal/model"
	"github.com/zyuanx/research-sys/internal/pkg/constant"
	"github.com/zyuanx/research-sys/internal/pkg/errors"
	"github.com/zyuanx/research-sys/internal/pkg/errors/ecode"
	"github.com/zyuanx/research-sys/internal/pkg/response"
)

func (c *Controller) ResearchList(ctx *gin.Context) {
	req := model.ResearchListReq{}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		err = errors.Wrap(err, ecode.ValidateErr, "参数错误")
		response.JSON(ctx, err, nil)
		return
	}
	var researches []model.Research
	page, size := req.Page, req.Size
	var total int64

	query := make(map[string]interface{})

	if err := c.service.ResearchList(&researches, page, size, &total, query); err != nil {
		err = errors.Wrap(err, ecode.RecordListErr, "获取数据失败")
		response.JSON(ctx, err, nil)
		return
	}
	results := make([]model.ResearchRes, len(researches))
	for i, research := range researches {
		results[i] = research.ToRes()
	}
	response.JSON(ctx, nil, gin.H{
		"page":    page,
		"size":    size,
		"results": researches,
		"total":   total,
	})
}

func (c *Controller) ResearchRetrieve(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		err = errors.Wrap(err, ecode.ValidateErr, "ID错误")
		response.JSON(ctx, err, nil)
		return
	}
	research := model.Research{}
	if err = c.service.ResearchRetrieve(&research, id); err != nil {
		err = errors.Wrap(err, ecode.RecordRetrieveErr, "未找到记录")
		response.JSON(ctx, err, nil)
		return
	}
	response.JSON(ctx, nil, research.ToRes())
}

func (c *Controller) ResearchCreate(ctx *gin.Context) {
	req := model.ResearchCreateReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		err = errors.Wrap(err, ecode.ValidateErr, "参数错误")
		response.JSON(ctx, err, nil)
		return
	}

	var err error
	userId, exist := ctx.Get(constant.UserID)
	if !exist {
		err = errors.Wrap(err, ecode.AuthTokenErr, "未登录")
		response.JSON(ctx, err, nil)
		return
	}
	research := model.Research{
		Title:       req.Title,
		Description: req.Description,
		StartAt:     req.StartAt,
		EndAt:       req.EndAt,
		Once:        *req.Once,
		Open:        *req.Open,
		PublisherID: uint(userId.(int64)),
	}
	var config, pattern, items []byte

	if config, err = json.Marshal(req.Config); err != nil {
		err = errors.Wrap(err, ecode.ValidateErr, "参数错误")
		response.JSON(ctx, err, nil)
		return
	}
	research.Config = string(config)
	if pattern, err = json.Marshal(req.Pattern); err != nil {
		err = errors.Wrap(err, ecode.ValidateErr, "参数错误")
		response.JSON(ctx, err, nil)
		return
	}
	research.Pattern = string(pattern)
	if items, err = json.Marshal(req.Items); err != nil {
		err = errors.Wrap(err, ecode.ValidateErr, "参数错误")
		response.JSON(ctx, err, nil)
		return
	}
	research.Items = string(items)

	if err := c.service.ResearchCreate(&research); err != nil {
		err = errors.Wrap(err, ecode.RecordCreateErr, "创建失败")
		response.JSON(ctx, err, nil)
		return
	}
	response.JSON(ctx, nil, research.ToRes())
}

func (c *Controller) ResearchUpdate(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		err = errors.Wrap(err, ecode.ValidateErr, "ID错误")
		response.JSON(ctx, err, nil)
		return
	}
	req := model.ResearchUpdateReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		err = errors.Wrap(err, ecode.ValidateErr, "参数错误")
		response.JSON(ctx, err, nil)
		return
	}
	research := model.Research{}
	if err = c.service.ResearchRetrieve(&research, id); err != nil {
		err = errors.Wrap(err, ecode.RecordRetrieveErr, "未找到记录")
		response.JSON(ctx, err, nil)
		return
	}
	payload := map[string]interface{}{
		"title":       req.Title,
		"description": req.Description,
		"startAt":     req.StartAt,
		"endAt":       req.EndAt,
		"once":        req.Once,
		"open":        req.Open,
	}
	var config, pattern, items []byte
	if config, err = json.Marshal(req.Config); err != nil {
		err = errors.Wrap(err, ecode.ValidateErr, "参数错误")
		response.JSON(ctx, err, nil)
		return
	}
	payload["config"] = string(config)
	if pattern, err = json.Marshal(req.Pattern); err != nil {
		err = errors.Wrap(err, ecode.ValidateErr, "参数错误")
		response.JSON(ctx, err, nil)
		return
	}
	payload["pattern"] = string(pattern)
	if items, err = json.Marshal(req.Items); err != nil {
		err = errors.Wrap(err, ecode.ValidateErr, "参数错误")
		response.JSON(ctx, err, nil)
		return
	}
	payload["items"] = string(items)
	if err = c.service.ResearchUpdate(&research, payload); err != nil {
		err = errors.Wrap(err, ecode.RecordUpdateErr, "更新失败")
		response.JSON(ctx, err, nil)
		return
	}
	response.JSON(ctx, nil, research.ToRes())
}

func (c *Controller) ResearchDelete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		err = errors.Wrap(err, ecode.ValidateErr, "ID错误")
		response.JSON(ctx, err, nil)
		return
	}
	if err = c.service.ResearchDelete(id); err != nil {
		err = errors.Wrap(err, ecode.RecordDeleteErr, "删除失败")
		response.JSON(ctx, err, nil)
		return
	}
	response.JSON(ctx, nil, gin.H{})
}

// func (c *Controller) DownloadExcel(ctx *gin.Context) {
// 	var id int
// 	var err error
// 	if id, err = strconv.Atoi(ctx.Param("id")); err != nil {
// 		log.Println(err.Error())
// 		util.Fail(ctx, gin.H{}, "ID错误")
// 		return
// 	}

// 	research := model.Research{}
// 	if err = researchServices.Retrieve(&research, id); err != nil {
// 		log.Println(err.Error())
// 		util.Fail(ctx, gin.H{}, "获取信息失败")
// 		return
// 	}

// 	var items []map[string]interface{}
// 	if err = json.Unmarshal([]byte(research.Items), &items); err != nil {
// 		log.Println(err.Error())
// 		util.Fail(ctx, gin.H{}, "数据解析失败")
// 		return
// 	}
// 	titles := []interface{}{"序号"}
// 	var fieldIDs []string
// 	for _, item := range items {
// 		titles = append(titles, item["label"].(string))
// 		fieldIDs = append(fieldIDs, item["fieldID"].(string))
// 	}
// 	titles = append(titles, "IP地址", "填写时间")

// 	var openRecords []model.OpenRecord
// 	if err = openRecordServices.ListByResearchID(&openRecords, research.ID); err != nil {
// 		log.Println(err.Error())
// 		util.Fail(ctx, nil, "获取数据失败")
// 		return
// 	}
// 	// 1. start generate excel
// 	xlsx := excelize.NewFile()
// 	// 2. new StreamWriter
// 	streamWriter, _ := xlsx.NewStreamWriter("Sheet1")
// 	if _, err = xlsx.NewStyle(`{"font":{"color":"#777777"}}`); err != nil {
// 		log.Println(err.Error())
// 		util.Fail(ctx, nil, "设置样式失败")
// 		return
// 	}
// 	// 3. write title
// 	if err = streamWriter.SetRow("A1", titles); err != nil {
// 		log.Println(err.Error())
// 		util.Fail(ctx, nil, "写入Row失败")
// 		return
// 	}

// 	// 4. write record data
// 	for k, v := range openRecords {
// 		row := []interface{}{k + 1}
// 		values := make(map[string]interface{})
// 		if err = json.Unmarshal([]byte(v.Values), &values); err != nil {
// 			log.Println(err.Error())
// 			util.Fail(ctx, gin.H{}, "数据解析失败")
// 			return
// 		}
// 		for _, fieldID := range fieldIDs {
// 			row = append(row, values[fieldID])
// 		}
// 		row = append(row, v.IPAddress)
// 		row = append(row, v.CreatedAt.Format("2006-01-02 15:04:05"))

// 		cell, _ := excelize.CoordinatesToCellName(1, k+2)
// 		if err = streamWriter.SetRow(cell, row); err != nil {
// 			log.Println(err.Error())
// 			util.Fail(ctx, nil, "写入Row失败")
// 			return
// 		}
// 	}
// 	// 5. flush streamWriter
// 	if err = streamWriter.Flush(); err != nil {
// 		log.Println(err.Error())
// 		util.Fail(ctx, nil, "刷新失败")
// 		return
// 	}
// 	ctx.Header("response-type", "blob")
// 	data, _ := xlsx.WriteToBuffer()
// 	ctx.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", data.Bytes())
// }
