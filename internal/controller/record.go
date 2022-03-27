package controller

import (
	"errors"
	"gin-research-sys/internal/form"
	"gin-research-sys/internal/model"
	"gin-research-sys/internal/service"
	"gin-research-sys/internal/util"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
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

//var recordMgoServices = service.NewRecordMgoService()

func (r RecordController) List(ctx *gin.Context) {
	pagination := form.Pagination{}
	if err := ctx.ShouldBindQuery(&pagination); err != nil {
		log.Println(err.Error())
		util.Success(ctx, nil, "query error")
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
	//idString := ctx.Param("id")
	//id, err := strconv.Atoi(idString)
	//if err != nil {
	//	log.Println(err.Error())
	//	util.Fail(ctx, gin.H{}, "param is error")
	//	return
	//}
	//record := model.Record{}
	//if err = recordServices.Retrieve(&record, id); err != nil {
	//	log.Println(err.Error())
	//	util.Fail(ctx, gin.H{}, "retrieve fail")
	//	return
	//}
	//recordMgo := model.RecordMgo{}
	//if err = recordMgoServices.Retrieve(&recordMgo, record.RecordID); err != nil {
	//	log.Println(err.Error())
	//	util.Fail(ctx, gin.H{}, "retrieve fail")
	//	return
	//}
	//
	//util.Success(ctx, gin.H{"record": gin.H{
	//	"id":          record.ID,
	//	"title":       record.Title,
	//	"recordID":    record.RecordID,
	//	"fieldsValue": recordMgo.FieldsValue,
	//}}, "")
}

func (r RecordController) Create(ctx *gin.Context) {
	//createForm := form.RecordCreateForm{}
	//if err := ctx.ShouldBindJSON(&createForm); err != nil {
	//	log.Println(err.Error())
	//	util.Fail(ctx, gin.H{}, "payload error")
	//	return
	//}
	//// there needs mongo transaction
	//recordMgo := model.RecordMgo{
	//	FieldsValue: createForm.FieldsValue,
	//}
	//result, err := recordMgoServices.Create(&recordMgo)
	//if err != nil {
	//	log.Println(err.Error())
	//	util.Fail(ctx, gin.H{}, "create error")
	//	return
	//}
	//claims := jwt.ExtractClaims(ctx)
	//id := int(claims["id"].(float64))
	//user := model.User{}
	//if err = userServices.Retrieve(&user, id); err != nil {
	//	util.Fail(ctx, gin.H{}, "record not found")
	//	return
	//}
	//
	//record := model.Record{
	//	Title:      createForm.Title,
	//	ResearchID: createForm.ResearchID,
	//	IP:         ctx.ClientIP(),
	//	UserID:     int(user.ID),
	//}
	//if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
	//	record.RecordID = oid.Hex()
	//} else {
	//	log.Println(ok)
	//	util.Fail(ctx, gin.H{}, "create fail")
	//	return
	//}
	//if err = recordServices.Create(&record); err != nil {
	//	log.Println(err.Error())
	//	util.Fail(ctx, gin.H{}, "create fail")
	//	return
	//}
	//util.Success(ctx, gin.H{}, "create success")
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
