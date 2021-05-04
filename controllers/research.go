package controllers

import (
	"gin-research-sys/models"
	"gin-research-sys/pkg/req"
	"gin-research-sys/pkg/res"
	"gin-research-sys/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

type IResearchController interface {
	List(ctx *gin.Context)
	Create(ctx *gin.Context)
	Retrieve(ctx *gin.Context)
	Update(ctx *gin.Context)
	Destroy(ctx *gin.Context)
}
type ResearchController struct {
}

func NewResearchController() ResearchController {
	return ResearchController{}
}

var researchServices = services.NewResearchListService()

func (r ResearchController) List(ctx *gin.Context) {
	pg := req.PaginationQuery{}
	if err := ctx.ShouldBindQuery(&pg); err != nil {
		res.Success(ctx, nil, err.Error())
		return
	}

	var total int64
	researches, err := researchServices.List(int64(pg.Page), int64(pg.Size), &total)
	if err != nil {
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
	research := models.ResearchList{}
	if err := ctx.ShouldBindJSON(&research); err != nil {
		log.Println(err.Error())
		res.Fail(ctx, gin.H{}, "payload error")
	}
	if err := researchServices.Create(&research); err != nil {
		log.Println(err.Error())
		res.Fail(ctx, gin.H{}, "create fail")
	}
	res.Success(ctx, gin.H{"research": research}, "create success")
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
