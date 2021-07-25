package controller

import (
	"gin-research-sys/internal/form"
	"gin-research-sys/internal/model"
	"gin-research-sys/internal/service"
	"gin-research-sys/internal/util"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

var permissionServices = service.NewPermissionService()

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
	pagination := form.Pagination{}
	if err := ctx.ShouldBindQuery(&pagination); err != nil {
		util.Fail(ctx, nil, err.Error())
		return
	}
	var permissions []model.Permission
	var total int64
	if err := permissionServices.List(pagination.Page, pagination.Size, &permissions, &total); err != nil {
		util.Success(ctx, nil, err.Error())
		return
	}
	util.Success(ctx, gin.H{
		"page":    pagination.Page,
		"size":    pagination.Size,
		"results": permissions,
		"total":   total,
	}, "")
}

func (p PermissionController) Retrieve(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "param is error")
		return
	}
	permission := model.Permission{}
	if err = permissionServices.Retrieve(&permission, id); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "retrieve fail")
		return
	}
	util.Success(ctx, gin.H{"permission": permission}, "retrieve success")
}

func (p PermissionController) Create(ctx *gin.Context) {
	createForm := form.PermissionCreateForm{}
	if err := ctx.ShouldBindJSON(&createForm); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "payload is error")
		return
	}
	permission := model.Permission{
		Group: createForm.Group,
		Path: createForm.Path,
		Method: createForm.Method,
		Desc: createForm.Desc,
		Index: createForm.Index,
	}

	if err := permissionServices.Create(&permission); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "create fail")
		return
	}
	util.Success(ctx, gin.H{}, "create success")
}
func (p PermissionController) Update(ctx *gin.Context) {

	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "param is error")
		return
	}
	//updateForm := form.PermissionUpdateForm{}
	//if err = ctx.ShouldBindJSON(&updateForm); err != nil {
	//	log.Println(err.Error())
	//	res.Fail(ctx, gin.H{}, "payload is error")
	//	return
	//}
	permission := model.Permission{}
	if err = permissionServices.Retrieve(&permission, id); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "retrieve fail")
		return
	}
	if err = ctx.ShouldBindJSON(&permission); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "payload is error")
		return
	}
	if err = permissionServices.Update(&permission); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "update fail")
		return
	}
	util.Success(ctx, gin.H{}, "update success")
}

func (p PermissionController) Destroy(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "param is error")
		return
	}
	if err = permissionServices.Destroy(id); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "delete fail")
		return
	}
	util.Success(ctx, gin.H{}, "delete success")
}
