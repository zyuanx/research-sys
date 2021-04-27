package controllers

import (
	"gin-research-sys/controllers/request"
	"gin-research-sys/controllers/response"
	"gin-research-sys/models"
	"gin-research-sys/services"
	"github.com/gin-gonic/gin"
)

type IRoleController interface {
	Create(ctx *gin.Context)
}
type RoleController struct {
}

func NewRoleController() RoleController {
	return RoleController{}
}

var roleServices = services.NewRoleService()

func (r RoleController) Create(ctx *gin.Context) {
	createRoleValidate := request.CreateRoleValidate{}
	if err := ctx.ShouldBindJSON(&createRoleValidate); err != nil {
		response.Fail(ctx, gin.H{}, err.Error())
		return
	}
	role := models.Role{
		Title: createRoleValidate.Title,
		Desc:  createRoleValidate.Desc,
	}
	err := roleServices.Create(&role)
	if err != nil {
		response.Fail(ctx, gin.H{}, err.Error())
		return
	}
	response.Success(ctx, gin.H{"role": role}, "")
}
