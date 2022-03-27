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

var roleServices = service.NewRoleService()

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
// @Summary list role
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param pagination query req.PaginationQuery true "PaginationQuery"
// @Success 200 {object} res.Result "成功后返回值"
// @Router /api/role [get]
func (r RoleController) List(ctx *gin.Context) {
	//pagination := form.Pagination{}
	//if err := ctx.ShouldBindQuery(&pagination); err != nil {
	//	util.Fail(ctx, nil, "query is error")
	//	return
	//}
	//var roles []model.Role
	//var total int64
	//if err := roleServices.List(pagination.Page, pagination.Size, &roles, &total); err != nil {
	//	util.Fail(ctx, nil, "list role error")
	//	return
	//}
	//util.Success(ctx, gin.H{
	//	"page":    pagination.Page,
	//	"size":    pagination.Size,
	//	"results": roles,
	//	"total":   total,
	//}, "")
}

// Create
// @Summary create role
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param role body req.RoleCreateReq true "角色信息"
// @Success 200 {object} res.Result "成功后返回值"
// @Router /api/role [post]
func (r RoleController) Create(ctx *gin.Context) {
	role := model.Role{}
	if err := ctx.ShouldBindJSON(&role); err != nil {
		util.Fail(ctx, gin.H{}, err.Error())
		return
	}
	//role := model.Role{}
	//utils.Struct2StructByJson(createRoleValidate, role)
	if err := roleServices.Create(&role); err != nil {
		util.Fail(ctx, gin.H{}, err.Error())
		return
	}
	util.Success(ctx, gin.H{"role": role}, "")
}

// Retrieve
// @Summary retrieve role
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param id path int true "ID"
// @Success 200 {object} res.Result "成功后返回值"
// @Router /api/role/{id} [get]
func (r RoleController) Retrieve(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		util.Fail(ctx, gin.H{}, err.Error())
		return
	}
	role := model.Role{}
	if err = roleServices.Retrieve(&role, id); err != nil {
		util.Fail(ctx, gin.H{}, err.Error())
		return
	}
	// get permission id list
	var permissions []int
	for _, value := range role.Permissions {
		permissions = append(permissions, int(value.ID))
	}
	util.Success(ctx, gin.H{"role": gin.H{
		"id":          role.ID,
		"title":       role.Title,
		"desc":        role.Desc,
		"permissions": permissions,
	}}, "")
}

// Update
// @Summary update role
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param id path int true "ID"
// @Param role body req.RoleUpdateReq true "角色信息"
// @Success 200 {object} res.Result "成功后返回值"
// @Router /api/role/{id} [put]
func (r RoleController) Update(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "param error")
		return
	}
	updateForm := form.RoleUpdateForm{}
	if err = ctx.ShouldBindJSON(&updateForm); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "payload error")
		return
	}
	role := model.Role{}
	if err = roleServices.Retrieve(&role, id); err != nil {
		util.Fail(ctx, gin.H{}, err.Error())
		return
	}
	role.Title = updateForm.Title
	role.Desc = updateForm.Desc
	if err = roleServices.Update(&role); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "update fail")
		return
	}
	if err = roleServices.UpdatePermission(&role, updateForm.Permissions); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "update fail")
		return
	}
	util.Success(ctx, gin.H{}, "update success")
}

// Destroy
// @Summary destroy role
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param id path int true "ID"
// @Success 200 {object} res.Result "成功后返回值"
// @Router /api/role/{id} [delete]
func (r RoleController) Destroy(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "param error")
		return
	}
	if err = roleServices.Destroy(id); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "delete fail")
		return
	}
	util.Success(ctx, gin.H{}, "delete success")
}
