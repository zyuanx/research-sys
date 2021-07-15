package controller

import (
	"gin-research-sys/internal/controller/req"
	"gin-research-sys/internal/controller/res"
	"gin-research-sys/internal/model"
	"gin-research-sys/internal/service"
	jwt "github.com/appleboy/gin-jwt/v2"
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

var recordServices = service.NewRecordService()
var recordMgoServices = service.NewRecordMgoService()

func (r RecordController) List(ctx *gin.Context) {
	pg := req.PaginationQuery{}
	if err := ctx.ShouldBindQuery(&pg); err != nil {
		log.Println(err.Error())
		res.Success(ctx, nil, "query error")
		return
	}

	var records []model.Record
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
	record := model.Record{}
	if err = recordServices.Retrieve(&record, id); err != nil {
		log.Println(err.Error())
		res.Fail(ctx, gin.H{}, "retrieve1 fail")
		return
	}
	recordMgo := model.RecordMgo{}
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
	recordMgo := model.RecordMgo{
		FieldsValue: createReq.FieldsValue,
	}
	result, err := recordMgoServices.Create(&recordMgo)
	if err != nil {
		log.Println(err.Error())
		res.Fail(ctx, gin.H{}, "create error")
		return
	}
	claims := jwt.ExtractClaims(ctx)
	id := int(claims["id"].(float64))
	user := model.User{}
	if err = userServices.Retrieve(&user, id); err != nil {
		res.Fail(ctx, gin.H{}, "record not found")
		return
	}

	record := model.Record{
		Title:      createReq.Title,
		ResearchID: createReq.ResearchID,
		IP:         ctx.ClientIP(),
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
