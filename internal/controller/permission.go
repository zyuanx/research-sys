package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zyuanx/research-sys/internal/form"
	"github.com/zyuanx/research-sys/internal/model"
	"github.com/zyuanx/research-sys/internal/pkg/errors"
	"github.com/zyuanx/research-sys/internal/pkg/errors/ecode"
	"github.com/zyuanx/research-sys/internal/pkg/response"
)

func (c *Controller) PermissionList(ctx *gin.Context) {
	pagination := form.Pagination{}
	if err := ctx.ShouldBindQuery(&pagination); err != nil {
		err = errors.Wrap(err, ecode.ValidateErr, "bind query error")
		response.JSON(ctx, err, nil)
		return
	}
	var permissions []model.Permission
	var total int64
	if err := c.service.PermissionList(pagination.Page, pagination.Size, &permissions, &total); err != nil {
		err = errors.Wrap(err, ecode.RecordListErr, "list error")
		response.JSON(ctx, err, nil)
		return
	}
	response.JSON(ctx, nil, gin.H{
		"page":    pagination.Page,
		"size":    pagination.Size,
		"results": permissions,
		"total":   total,
	})
}

func (c *Controller) PermissionRetrieve(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		err = errors.Wrap(err, ecode.ValidateErr, "param is error")
		response.JSON(ctx, err, nil)
		return
	}
	permission := model.Permission{}
	if err = c.service.PermissionRetrieve(&permission, id); err != nil {
		err = errors.Wrap(err, ecode.RecordRetrieveErr, "retrieve error")
		response.JSON(ctx, err, nil)
		return
	}
	response.JSON(ctx, nil, permission)
}

func (c *Controller) PermissionCreate(ctx *gin.Context) {
	createForm := form.PermissionCreateForm{}
	if err := ctx.ShouldBindJSON(&createForm); err != nil {
		err = errors.Wrap(err, ecode.ValidateErr, "payload is error")
		response.JSON(ctx, err, nil)
		return
	}
	permission := model.Permission{
		Group:  createForm.Group,
		Path:   createForm.Path,
		Method: createForm.Method,
		Desc:   createForm.Desc,
		Index:  createForm.Index,
	}

	if err := c.service.PermissionCreate(&permission); err != nil {
		err = errors.Wrap(err, ecode.RecordCreateErr, "create error")
		response.JSON(ctx, err, nil)
		return
	}
	response.JSON(ctx, nil, permission)
}
func (c *Controller) PermissionUpdate(ctx *gin.Context) {

	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		err = errors.Wrap(err, ecode.ValidateErr, "param is error")
		response.JSON(ctx, err, nil)
		return
	}
	//updateForm := form.PermissionUpdateForm{}
	//if err = ctx.ShouldBindJSON(&updateForm); err != nil {
	//	log.Println(err.Error())
	//	res.Fail(ctx, gin.H{}, "payload is error")
	//	return
	//}
	permission := model.Permission{}
	if err = c.service.PermissionRetrieve(&permission, id); err != nil {
		err = errors.Wrap(err, ecode.RecordRetrieveErr, "retrieve error")
		response.JSON(ctx, err, nil)
		return
	}
	if err = ctx.ShouldBindJSON(&permission); err != nil {
		err = errors.Wrap(err, ecode.ValidateErr, "payload is error")
		response.JSON(ctx, err, nil)
		return
	}
	if err = c.service.PermissionUpdate(&permission); err != nil {
		err = errors.Wrap(err, ecode.RecordUpdateErr, "update error")
		response.JSON(ctx, err, nil)
		return
	}
	response.JSON(ctx, nil, permission)
}

func (c *Controller) PermissionDelete(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		err = errors.Wrap(err, ecode.ValidateErr, "param is error")
		response.JSON(ctx, err, nil)
		return
	}
	if err = c.service.PermissionDestroy(id); err != nil {
		err = errors.Wrap(err, ecode.RecordDeleteErr, "delete error")
		response.JSON(ctx, err, nil)
		return
	}
	response.JSON(ctx, nil, nil)
}
