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

type IResearchController interface {
	List(ctx *gin.Context)
	Create(ctx *gin.Context)
	Retrieve(ctx *gin.Context)
	Update(ctx *gin.Context)
	Destroy(ctx *gin.Context)

	UpdateStatus(ctx *gin.Context)
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
		res.Fail(ctx, nil, err.Error())
		return
	}

	var researches []models.Research
	var total int64
	if err := researchServices.List(pg.Page, pg.Size, &researches, &total); err != nil {
		res.Success(ctx, nil, err.Error())
		return
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
	research:= models.Research{}
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
func (r ResearchController) UpdateStatus(ctx *gin.Context) {
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
	research:= models.Research{}
	if err = researchServices.Retrieve(&research, id); err != nil {
		log.Println(err.Error())
		res.Fail(ctx, gin.H{}, "retrieve fail")
		return
	}
	research.Status = updateReq.Status
	if err = researchServices.Update(&research); err != nil {
		log.Println(err.Error())
		res.Fail(ctx, gin.H{}, "update fail")
		return
	}
	res.Success(ctx, gin.H{}, "update success")

}
