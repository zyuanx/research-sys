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

func (c Controller) RoleList(ctx *gin.Context) {
	pagination := form.Pagination{}
	if err := ctx.ShouldBindQuery(&pagination); err != nil {
		err = errors.Wrap(err, ecode.ValidateErr, "参数错误")
		response.JSON(ctx, err, nil)
		return
	}
	var roles []model.Role
	page, size := pagination.Page, pagination.Size
	var total int64
	query := make(map[string]interface{})
	if err := c.service.RoleList(&roles, page, size, &total, query); err != nil {
		err = errors.Wrap(err, ecode.ValidateErr, "获取角色失败")
		response.JSON(ctx, err, nil)
		return
	}
	response.JSON(ctx, nil, gin.H{
		"page":    pagination.Page,
		"size":    pagination.Size,
		"results": roles,
		"total":   total,
	})
}

func (c Controller) RoleRetrieve(ctx *gin.Context) {
	var id int
	var err error
	if id, err = strconv.Atoi(ctx.Param("id")); err != nil {
		err = errors.Wrap(err, ecode.ValidateErr, "ID错误")
		response.JSON(ctx, err, nil)
		return
	}
	role := &model.Role{}
	if err = c.service.RoleRetrieve(role, id); err != nil {
		err = errors.Wrap(err, ecode.RecordRetrieveErr, "未找到记录")
		response.JSON(ctx, err, nil)
		return
	}
	response.JSON(ctx, nil, role)
}

func (c Controller) RoleCreate(ctx *gin.Context) {
	role := &model.Role{}
	if err := ctx.ShouldBindJSON(role); err != nil {
		err = errors.Wrap(err, ecode.ValidateErr, "参数错误")
		response.JSON(ctx, err, nil)
		return
	}
	if err := c.service.RoleCreate(role); err != nil {
		err = errors.Wrap(err, ecode.RecordCreateErr, "创建角色失败")
		response.JSON(ctx, err, nil)
		return
	}
	response.JSON(ctx, nil, role)
}

func (c Controller) RoleUpdate(ctx *gin.Context) {
	var id int
	var err error
	id, err = strconv.Atoi(ctx.Param("id"))
	if err != nil {
		err = errors.Wrap(err, ecode.ValidateErr, "ID错误")
		response.JSON(ctx, err, nil)
		return
	}
	req := model.RoleUpdateReq{}
	if err = ctx.ShouldBindJSON(&req); err != nil {
		err = errors.Wrap(err, ecode.ValidateErr, "参数错误")
		response.JSON(ctx, err, nil)
		return
	}
	role := &model.Role{}
	if err = c.service.RoleRetrieve(role, id); err != nil {
		err = errors.Wrap(err, ecode.RecordRetrieveErr, "未找到记录")
		response.JSON(ctx, err, nil)
		return
	}
	payload := map[string]interface{}{
		"title": req.Title,
		"desc":  req.Desc,
	}
	if err = c.service.RoleUpdate(role, payload); err != nil {
		err = errors.Wrap(err, ecode.RecordUpdateErr, "更新角色失败")
		response.JSON(ctx, err, nil)
		return
	}
	response.JSON(ctx, nil, role)
}

func (c Controller) RoleDelete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		err = errors.Wrap(err, ecode.ValidateErr, "ID错误")
		response.JSON(ctx, err, nil)
		return
	}
	if err = c.service.RoleDelete(id); err != nil {
		err = errors.Wrap(err, ecode.RecordDeleteErr, "删除角色失败")
		response.JSON(ctx, err, nil)
		return
	}
	response.JSON(ctx, nil, nil)
}
