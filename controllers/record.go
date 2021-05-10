package controllers

import (
	"gin-research-sys/controllers/req"
	"gin-research-sys/controllers/res"
	"gin-research-sys/middlewares"
	"gin-research-sys/models"
	"gin-research-sys/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"strconv"
)

type IRecordController interface {
	List(ctx *gin.Context)
	Create(ctx *gin.Context)
	Retrieve(ctx *gin.Context)
}
type RecordController struct{}

func NewRecordController() IRecordController {
	return RecordController{}
}

var recordServices = services.NewRecordService()
var recordMgoServices = services.NewRecordMgoService()

func (r RecordController) List(ctx *gin.Context) {
	pg := req.PaginationQuery{}
	if err := ctx.ShouldBindQuery(&pg); err != nil {
		log.Println(err.Error())
		res.Success(ctx, nil, "query error")
		return
	}

	var records []models.Record
	var total int64
	if err := recordServices.List(pg.Page, pg.Size, &records, &total); err != nil {
		res.Success(ctx, nil, err.Error())
		return
	}
	res.Success(ctx, gin.H{
		"page":    pg.Page,
		"size":    pg.Size,
		"results": records,
		"total":   total,
	}, "")
}

func (r RecordController) Retrieve(ctx *gin.Context) {

	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Println(err.Error())
		res.Fail(ctx, gin.H{}, "param is error")
		return
	}
	record := models.Record{}
	if err = recordServices.Retrieve(&record, id); err != nil {
		log.Println(err.Error())
		res.Fail(ctx, gin.H{}, "retrieve1 fail")
		return
	}
	recordMgo := models.RecordMgo{}
	if err = recordMgoServices.Retrieve(&recordMgo, record.RecordID); err != nil {
		log.Println(err.Error())
		res.Fail(ctx, gin.H{}, "retrieve2 fail")
		return
	}

	res.Success(ctx, gin.H{"record": gin.H{
		"id":          record.ID,
		"title":       record.Title,
		"recordID":    record.RecordID,
		"fieldsValue": recordMgo.FieldsValue,
	}}, "")
}

func (r RecordController) Create(ctx *gin.Context) {
	createReq := req.RecordCreateReq{}
	if err := ctx.ShouldBindJSON(&createReq); err != nil {
		log.Println(err.Error())
		res.Fail(ctx, gin.H{}, "payload error")
		return
	}
	// there needs mongo transaction
	recordMgo := models.RecordMgo{
		FieldsValue: createReq.FieldsValue,
	}
	result, err := recordMgoServices.Create(&recordMgo)
	if err != nil {
		log.Println(err.Error())
		res.Fail(ctx, gin.H{}, "create error")
		return
	}
	user := middlewares.JWTAuthMiddleware.IdentityHandler(ctx).(models.User)
	record := models.Record{
		Title:      createReq.Title,
		ResearchID: createReq.ResearchID,
		UserID:     int(user.ID),
	}
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		record.RecordID = oid.Hex()
	} else {
		log.Println(ok)
		res.Fail(ctx, gin.H{}, "create fail")
		return
	}
	if err = recordServices.Create(&record); err != nil {
		log.Println(err.Error())
		res.Fail(ctx, gin.H{}, "create fail")
		return
	}
	res.Success(ctx, gin.H{}, "create success")
}
