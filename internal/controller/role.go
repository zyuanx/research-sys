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
	Retrieve(ctx *gin.Context)
	Create(ctx *gin.Context)
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
	pagination := form.Pagination{}
	if err := ctx.ShouldBindQuery(&pagination); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, nil, "参数错误")
		return
	}
	var roles []model.Role
	page, size := pagination.Page, pagination.Size
	var total int64
	query := make(map[string]interface{})
	if err := roleServices.List(&roles, page, size, &total, query); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, nil, "获取数据失败")
		return
	}
	util.Success(ctx, gin.H{
		"page":    pagination.Page,
		"size":    pagination.Size,
		"results": roles,
		"total":   total,
	}, "")
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
	var id int
	var err error
	if id, err = strconv.Atoi(ctx.Param("id")); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "ID错误")
		return
	}
	role := model.Role{}
	if err = roleServices.Retrieve(&role, id); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "未找到记录")
		return
	}
	util.Success(ctx, gin.H{"role": role}, "获取角色信息成功")
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
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "参数错误")
		return
	}
	//role := model.Role{}
	//utils.Struct2StructByJson(createRoleValidate, role)
	if err := roleServices.Create(&role); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "创建失败")
		return
	}
	util.Success(ctx, gin.H{}, "创建成功")
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
	var id int
	var err error
	id, err = strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "ID错误")
		return
	}
	updateForm := form.RoleUpdateForm{}
	if err = ctx.ShouldBindJSON(&updateForm); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "参数错误")
		return
	}
	role := model.Role{}
	if err = roleServices.Retrieve(&role, id); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "获取记录失败")
		return
	}
	payload := map[string]interface{}{
		"title": updateForm.Title,
		"desc":  updateForm.Desc,
	}
	if err = roleServices.Update(&role, payload); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "更新角色失败")
		return
	}
	if err = roleServices.UpdatePermission(&role, updateForm.Permissions); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "更新角色失败")
		return
	}
	util.Success(ctx, gin.H{}, "更新角色成功")
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
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "参数错误")
		return
	}
	if err = roleServices.Destroy(id); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "删除角色失败")
		return
	}
	util.Success(ctx, gin.H{}, "删除角色成功")
}
