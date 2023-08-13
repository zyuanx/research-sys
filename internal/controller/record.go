package controller

// import (
// 	"errors"
// 	"gin-research-sys/internal/util"

// 	"log"
// 	"strconv"
// 	"time"

// 	jwt "github.com/appleboy/gin-jwt/v2"
// 	"github.com/gin-gonic/gin"
// 	"github.com/zyuanx/research-sys/internal/form"
// 	"github.com/zyuanx/research-sys/internal/model"
// 	"github.com/zyuanx/research-sys/internal/request"
// 	"gorm.io/gorm"
// )

// func (c *Controller) List(ctx *gin.Context) {
// 	pagination := form.Pagination{}
// 	if err := ctx.ShouldBindQuery(&pagination); err != nil {
// 		log.Println(err.Error())
// 		util.Fail(ctx, nil, "参数错误")
// 		return
// 	}

// 	var records []model.Record
// 	var total int64

// 	// not admin, retrieve created by yourself
// 	query := make(map[string]interface{})
// 	claims := jwt.ExtractClaims(ctx)
// 	id := int(claims["id"].(float64))
// 	username := claims["username"].(string)
// 	if username != "admin" {
// 		query["publisher"] = id
// 	}
// 	if err := c.service.RecordList(pagination.Page, pagination.Size, &records, &total, query); err != nil {
// 		util.Success(ctx, nil, err.Error())
// 		return
// 	}
// 	util.Success(ctx, gin.H{
// 		"page":    pagination.Page,
// 		"size":    pagination.Size,
// 		"results": records,
// 		"total":   total,
// 	}, "")
// }

// func (c *Controller) Retrieve(ctx *gin.Context) {
// 	idString := ctx.Param("id")
// 	id, err := strconv.Atoi(idString)
// 	if err != nil {
// 		log.Println(err.Error())
// 		util.Fail(ctx, gin.H{}, "参数错误")
// 		return
// 	}
// 	record := model.Record{}
// 	if err = recordServices.Retrieve(&record, id); err != nil {
// 		log.Println(err.Error())
// 		util.Fail(ctx, gin.H{}, "获取记录失败")
// 		return
// 	}

// 	util.Success(ctx, gin.H{"record": record}, "获取成功")
// }

// func (c *Controller) Create(ctx *gin.Context) {
// 	createPayload := request.RecordCreatePayload{}
// 	if err := ctx.ShouldBindJSON(&createPayload); err != nil {
// 		log.Println(err.Error())
// 		util.Fail(ctx, gin.H{}, "参数错误")
// 		return
// 	}

// 	// 获取用户信息
// 	claims := jwt.ExtractClaims(ctx)
// 	userID := int(claims["id"].(float64))

// 	// 获取问卷信息
// 	research := model.Research{}
// 	if err := researchServices.Retrieve(&research, createPayload.ResearchID); err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			util.Fail(ctx, gin.H{}, "未找到记录")
// 			return
// 		}
// 		log.Println(err.Error())
// 		util.Fail(ctx, gin.H{}, "获取数据失败")
// 		return
// 	}
// 	// 检查是否在时间范围内
// 	now := time.Now()
// 	if now.Before(research.StartAt) || now.After(research.EndAt) {
// 		util.Fail(ctx, gin.H{}, "不在时间范围内")
// 		return
// 	}

// 	record := model.Record{
// 		ResearchID: createPayload.ResearchID,
// 		Values:     createPayload.Values,
// 		IPAddress:  ctx.ClientIP(),
// 		UserID:     userID,
// 	}
// 	if err := recordServices.Create(&record, &research); err != nil {
// 		log.Println(err.Error())
// 		util.Fail(ctx, gin.H{}, "创建失败")
// 		return
// 	}
// 	util.Success(ctx, gin.H{}, "创建成功")
// }

// func (c *Controller) Filled(ctx *gin.Context) {
// 	researchId := ctx.Param("id")

// 	claims := jwt.ExtractClaims(ctx)
// 	id := int(claims["id"].(float64))

// 	record := model.Record{}
// 	if err := recordServices.FindByResearchID(&record, researchId, id); err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			util.Success(ctx, gin.H{}, "")
// 			return
// 		}
// 		log.Println(err.Error())
// 		util.Fail(ctx, gin.H{}, "retrieve fail")
// 		return
// 	}
// 	util.Fail(ctx, gin.H{}, "already filled in")
// }
