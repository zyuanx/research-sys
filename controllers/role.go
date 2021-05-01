package controllers

import (
	"gin-research-sys/controllers/request"
	"gin-research-sys/controllers/response"
	"gin-research-sys/models"
	"gin-research-sys/services"
	"github.com/gin-gonic/gin"
)

var roleServices = services.NewRoleService()

type RoleController struct {
}

func NewRoleController() RoleController {
	return RoleController{}
}

type IRoleController interface {
	List(ctx *gin.Context)
	Create(ctx *gin.Context)
}

// Create
// @Summary create a new role
// @Description get string by ID
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param role body request.CreateRoleValidate true "角色"
// @Success 200 {object} response.Result "成功后返回值"
// @Router /api/role [post]
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
	if err := roleServices.Create(&role); err != nil {
		response.Fail(ctx, gin.H{}, err.Error())
		return
	}
	response.Success(ctx, gin.H{"role": role}, "")
}
