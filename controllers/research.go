package controllers

import (
	"gin-research-sys/controllers/req"
	"gin-research-sys/controllers/res"
	"gin-research-sys/middlewares"
	"gin-research-sys/models"
	"gin-research-sys/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type IResearchController interface {
	List(ctx *gin.Context)
	Create(ctx *gin.Context)
	Retrieve(ctx *gin.Context)
	Update(ctx *gin.Context)
	Destroy(ctx *gin.Context)
}
type ResearchController struct{}

func NewResearchController() ResearchController {
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
	research := bson.M{}
	id := ctx.Param("id")
	if err := researchServices.Retrieve(&research, id); err != nil {
		res.Fail(ctx, gin.H{}, err.Error())
		return
	}
	res.Success(ctx, gin.H{"research": research}, "")
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
		research.ResearchID = oid.String()
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
	research := bson.M{}
	if err := ctx.ShouldBindJSON(&research); err != nil {
		res.Fail(ctx, gin.H{}, err.Error())
		return
	}
	if err := researchServices.Update(idString, research); err != nil {
		res.Fail(ctx, gin.H{}, "update fail")
		log.Println(err.Error())
		return
	}
	res.Success(ctx, gin.H{}, "update success")
}

func (r ResearchController) Destroy(ctx *gin.Context) {
	panic("implement me")
}
