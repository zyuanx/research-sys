package controllers

import (
	"gin-research-sys/controllers/res"
	"gin-research-sys/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type IResearchController interface {
	Retrieve(ctx *gin.Context)
}
type ResearchController struct {
}

func NewResearchController() ResearchController {
	return ResearchController{}
}

var researchServices = services.NewResearchListService()

func (r ResearchController) Retrieve(ctx *gin.Context) {
	research := bson.M{}
	id := ctx.Param("id")
	if err := researchServices.Retrieve(&research, id); err != nil {
		res.Fail(ctx, gin.H{}, err.Error())
		return
	}
	res.Success(ctx, gin.H{"research": research}, "")
}
