package controllers

import (
	"gin-research-sys/models"
	"gin-research-sys/pkg/req"
	"gin-research-sys/pkg/res"
	"gin-research-sys/services"
	"github.com/gin-gonic/gin"
	"strconv"
)

var permissionServices = services.NewPermissionService()

type IPermissionController interface {
	List(ctx *gin.Context)
	Create(ctx *gin.Context)
	Retrieve(ctx *gin.Context)
	Update(ctx *gin.Context)
	Destroy(ctx *gin.Context)
}
type PermissionController struct{}

func NewPermissionController() PermissionController {
	return PermissionController{}
}

func (p PermissionController) List(ctx *gin.Context) {
	pg := req.PaginationQuery{}
	if err := ctx.ShouldBindQuery(&pg); err != nil {
		res.Success(ctx, nil, err.Error())
		return
	}
	var permissions []models.Permission
	var total int64
	if err := permissionServices.List(pg.Page, pg.Size, &permissions, &total); err != nil {
		res.Success(ctx, nil, err.Error())
		return
	}
	res.Success(ctx, gin.H{
		"page":    pg.Page,
		"size":    pg.Size,
		"results": permissions,
		"total":   total,
	}, "")
}

func (p PermissionController) Create(ctx *gin.Context) {
	permission := models.Permission{}
	if err := ctx.ShouldBindJSON(&permission); err != nil {
		res.Fail(ctx, gin.H{}, err.Error())
	}

	if err := permissionServices.Create(&permission); err != nil {
		res.Fail(ctx, gin.H{}, err.Error())
	}
	res.Success(ctx, gin.H{"role": permission}, "")
}

func (p PermissionController) Retrieve(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		res.Fail(ctx, gin.H{}, err.Error())
	}
	permission := models.Permission{}
	if err = permissionServices.Retrieve(&permission, id); err != nil {
		res.Fail(ctx, gin.H{}, err.Error())
		return
	}
	res.Success(ctx, gin.H{"permission": permission}, "")
}
