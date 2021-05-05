package controllers

import (
	"gin-research-sys/controllers/req"
	"gin-research-sys/controllers/res"
	"gin-research-sys/models"
	"gin-research-sys/services"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

var permissionServices = services.NewPermissionService()

type IPermissionController interface {
	List(ctx *gin.Context)
	Retrieve(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Destroy(ctx *gin.Context)
}
type PermissionController struct{}

func NewPermissionController() IPermissionController {
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


func (p PermissionController) Retrieve(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Println(err.Error())
		res.Fail(ctx, gin.H{}, "param is error")
		return
	}
	permission := models.Permission{}
	if err = permissionServices.Retrieve(&permission, id); err != nil {
		log.Println(err.Error())
		res.Fail(ctx, gin.H{}, "retrieve fail")
		return
	}
	res.Success(ctx, gin.H{"permission": permission}, "retrieve success")
}

func (p PermissionController) Create(ctx *gin.Context) {
	permission := models.Permission{}
	if err := ctx.ShouldBindJSON(&permission); err != nil {
		log.Println(err.Error())
		res.Fail(ctx, gin.H{}, "payload is error")
		return
	}
	if err := permissionServices.Create(&permission); err != nil {
		log.Println(err.Error())
		res.Fail(ctx, gin.H{}, "create fail")
		return
	}
	res.Success(ctx, gin.H{}, "create success")
}
func (p PermissionController) Update(ctx *gin.Context) {
	panic("implement me")
}

func (p PermissionController) Destroy(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Println(err.Error())
		res.Fail(ctx, gin.H{}, "param error")
		return
	}
	if err = permissionServices.Destroy(id); err != nil {
		log.Println(err.Error())
		res.Fail(ctx, gin.H{}, "delete fail")
		return
	}
	res.Success(ctx, gin.H{}, "delete success")
}
