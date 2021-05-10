package controllers

import (
	"gin-research-sys/controllers/req"
	"gin-research-sys/controllers/res"
	"gin-research-sys/middlewares"
	"gin-research-sys/models"
	"gin-research-sys/services"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"strconv"
)

type IResearchController interface {
	List(ctx *gin.Context)
	Create(ctx *gin.Context)
	Retrieve(ctx *gin.Context)
	Update(ctx *gin.Context)
	Destroy(ctx *gin.Context)

	DownloadExcel(ctx *gin.Context)
}
type ResearchController struct{}

func NewResearchController() IResearchController {
	return ResearchController{}
}

var researchServices = services.NewResearchService()
var researchMgoServices = services.NewResearchMgoService()

func (r ResearchController) List(ctx *gin.Context) {
	pg := req.PaginationQuery{}
	if err := ctx.ShouldBindQuery(&pg); err != nil {
		log.Println(err.Error())
		res.Success(ctx, nil, "query error")
		return
	}

	var researches []models.Research
	var total int64
	if err := researchServices.List(pg.Page, pg.Size, &researches, &total); err != nil {
		log.Println(err.Error())
		res.Success(ctx, nil, "list error")
	}
	res.Success(ctx, gin.H{
		"page":    pg.Page,
		"size":    pg.Size,
		"results": researches,
		"total":   total,
	}, "")
}

func (r ResearchController) Retrieve(ctx *gin.Context) {

	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Println(err.Error())
		res.Fail(ctx, gin.H{}, "param is error")
		return
	}
	research := models.Research{}
	if err = researchServices.Retrieve(&research, id); err != nil {
		log.Println(err.Error())
		res.Fail(ctx, gin.H{}, "retrieve1 fail")
		return
	}
	researchMgo := models.ResearchMgo{}
	if err = researchMgoServices.Retrieve(&researchMgo, research.ResearchID); err != nil {
		log.Println(err.Error())
		res.Fail(ctx, gin.H{}, "retrieve2 fail")
		return
	}

	res.Success(ctx, gin.H{"research": gin.H{
		"id":          research.ID,
		"title":       research.Title,
		"desc":        research.Desc,
		"status":      research.Status,
		"once":        research.Once,
		"researchID":  research.ResearchID,
		"fieldsValue": researchMgo.FieldsValue,
		"detail":      researchMgo.Detail,
		"rules":       researchMgo.Rules,
	}}, "")
}

func (r ResearchController) Create(ctx *gin.Context) {
	createReq := req.ResearchCreateReq{}
	if err := ctx.ShouldBindJSON(&createReq); err != nil {
		log.Println(err.Error())
		res.Fail(ctx, gin.H{}, "payload error")
		return
	}
	// there needs mongo transaction
	researchMgo := models.ResearchMgo{
		FieldsValue: createReq.FieldsValue,
		Detail:      createReq.Detail,
		Rules:       createReq.Rules,
	}
	result, err := researchMgoServices.Create(&researchMgo)
	if err != nil {
		log.Println(err.Error())
		res.Fail(ctx, gin.H{}, "create error")
		return
	}
	user := middlewares.JWTAuthMiddleware.IdentityHandler(ctx).(models.User)
	research := models.Research{
		Title:  createReq.Title,
		Desc:   createReq.Desc,
		Once:   createReq.Once,
		UserID: int(user.ID),
	}
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		research.ResearchID = oid.Hex()
	} else {
		log.Println(ok)
		res.Fail(ctx, gin.H{}, "create fail")
		return
	}
	if err = researchServices.Create(&research); err != nil {
		log.Println(err.Error())
		res.Fail(ctx, gin.H{}, "create fail")
		return
	}
	res.Success(ctx, gin.H{}, "create success")
}

func (r ResearchController) Update(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Println(err.Error())
		res.Fail(ctx, gin.H{}, "param is error")
		return
	}
	updateReq := req.ResearchUpdateReq{}
	if err = ctx.ShouldBindJSON(&updateReq); err != nil {
		log.Println(err.Error())
		res.Fail(ctx, gin.H{}, "payload is error")
		return
	}
	research := models.Research{}
	if err = researchServices.Retrieve(&research, id); err != nil {
		log.Println(err.Error())
		res.Fail(ctx, gin.H{}, "retrieve fail")
		return
	}
	research.Title = updateReq.Title
	research.Desc = updateReq.Desc
	research.Once = updateReq.Once
	research.Status = updateReq.Status
	if err = researchServices.Update(&research); err != nil {
		log.Println(err.Error())
		res.Fail(ctx, gin.H{}, "update fail")
		return
	}
	//res.Success(ctx, gin.H{}, "update success")
	//research := bson.M{}
	//if err := researchServices.Update(idString, research); err != nil {
	//	res.Fail(ctx, gin.H{}, "update fail")
	//	log.Println(err.Error())
	//	return
	//}
	res.Success(ctx, gin.H{}, "update success")
}

func (r ResearchController) Destroy(ctx *gin.Context) {
	panic("implement me")
}

func (r ResearchController) DownloadExcel(ctx *gin.Context) {
	idString := ctx.Param("id")
	// get research
	researchMgo := models.ResearchMgo{}
	if err := researchMgoServices.Retrieve(&researchMgo, idString); err != nil {
		log.Println(err.Error())
		res.Fail(ctx, gin.H{}, "retrieve2 fail")
		return
	}
	// get research field
	var fields []string
	titleRow := make([]interface{}, len(researchMgo.Detail))
	for i, v := range researchMgo.Detail {
		titleRow[i] = v["label"].(string)
		fields = append(fields, v["fieldId"].(string))
	}

	// get record
	var records []models.Record
	var total int64
	if err := recordServices.ListID(idString, &records, &total); err != nil {
		log.Println(err.Error())
		res.Success(ctx, nil, "list error")
		return
	}
	xlsx := excelize.NewFile()
	streamWriter, err := xlsx.NewStreamWriter("Sheet1")
	if err != nil {
		println(err.Error())
	}
	if err := streamWriter.SetRow("A1", titleRow);err != nil {
		return
	}
	if err := streamWriter.Flush(); err != nil {
		println(err.Error())
	}

	//for colID := 0; colID < 50; colID++ {
	//	row[colID] = rand.Intn(640000)
	//}

	if err := xlsx.SetCellValue("Sheet1", "A2", "asdas"); err != nil {
		return
	}
	ctx.Header("response-type", "blob")
	data, _ := xlsx.WriteToBuffer()
	ctx.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", data.Bytes())
}
