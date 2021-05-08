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

var roleServices = services.NewRoleService()

type IRoleController interface {
	List(ctx *gin.Context)
	Create(ctx *gin.Context)
	Retrieve(ctx *gin.Context)
	Update(ctx *gin.Context)
	Destroy(ctx *gin.Context)
}
type RoleController struct{}

func NewRoleController() IRoleController {
	return RoleController{}
}

// List
// @Summary list all role
// @Description
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param pagination query req.PaginationQuery true "PaginationQuery"
// @Success 200 {object} res.Result "成功后返回值"
// @Failure 400 {object} res.Fail
// @Router /api/role [get]
func (r RoleController) List(ctx *gin.Context) {
	pg := req.PaginationQuery{}
	if err := ctx.ShouldBindQuery(&pg); err != nil {
		res.Fail(ctx, nil, err.Error())
		return
	}
	var roles []models.Role
	var total int64
	if err := roleServices.List(pg.Page, pg.Size, &roles, &total); err != nil {
		res.Success(ctx, nil, err.Error())
		return
	}
	res.Success(ctx, gin.H{
		"page":    pg.Page,
		"size":    pg.Size,
		"results": roles,
		"total":   total,
	}, "")
}

// Create
// @Summary create a new role
// @Description get string by ID
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param role body req.CreateRoleValidate true "角色"
// @Success 200 {object} res.Result "成功后返回值"
// @Router /api/role [post]
func (r RoleController) Create(ctx *gin.Context) {
	role := models.Role{}
	if err := ctx.ShouldBindJSON(&role); err != nil {
		res.Fail(ctx, gin.H{}, err.Error())
	}
	//role := models.Role{}
	//utils.Struct2StructByJson(createRoleValidate, role)
	if err := roleServices.Create(&role); err != nil {
		res.Fail(ctx, gin.H{}, err.Error())
	}
	res.Success(ctx, gin.H{"role": role}, "")
}

func (r RoleController) Retrieve(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		res.Fail(ctx, gin.H{}, err.Error())
	}
	role := models.Role{}
	if err = roleServices.Retrieve(&role, id); err != nil {
		res.Fail(ctx, gin.H{}, err.Error())
		return
	}
	var permissions []int
	if err := roleServices.ListPermission(&role, &permissions); err != nil {
		log.Println(err.Error())
		res.Fail(ctx, gin.H{}, "get roles error")
	}

	res.Success(ctx, gin.H{"role": gin.H{
		"id":          role.ID,
		"title":       role.Title,
		"desc":        role.Desc,
		"permissions": permissions,
	}}, "")
}

func (r RoleController) Update(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Println(err.Error())
		res.Fail(ctx, gin.H{}, "param error")
		return
	}
	updateReq := req.RoleUpdateReq{}
	if err = ctx.ShouldBindJSON(&updateReq); err != nil {
		log.Println(err.Error())
		res.Fail(ctx, gin.H{}, "payload error")
		return
	}
	role := models.Role{}
	if err = roleServices.Retrieve(&role, id); err != nil {
		res.Fail(ctx, gin.H{}, err.Error())
		return
	}
	role.Title = updateReq.Title
	role.Desc = updateReq.Desc
	if err = roleServices.Update(&role); err != nil {
		log.Println(err.Error())
		res.Fail(ctx, gin.H{}, "update fail")
		return
	}
	if err = roleServices.UpdatePermission(&role, updateReq.Permissions); err != nil {
		log.Println(err.Error())
		res.Fail(ctx, gin.H{}, "update fail")
		return
	}
	res.Success(ctx, gin.H{}, "update success")
}

func (r RoleController) Destroy(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Println(err.Error())
		res.Fail(ctx, gin.H{}, "param error")
		return
	}
	if err = roleServices.Destroy(id); err != nil {
		log.Println(err.Error())
		res.Fail(ctx, gin.H{}, "delete fail")
		return
	}
	res.Success(ctx, gin.H{}, "delete success")
}
