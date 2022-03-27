package controller

import (
	"gin-research-sys/internal/form"
	"gin-research-sys/internal/middleware"
	"gin-research-sys/internal/model"
	"gin-research-sys/internal/service"
	"gin-research-sys/internal/util"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strconv"
)

type IUserController interface {
	GetInfo(ctx *gin.Context)
	ChangePassword(ctx *gin.Context)

	List(ctx *gin.Context)
	Retrieve(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Destroy(ctx *gin.Context)
}

type UserController struct{}

func NewUserController() IUserController {
	return UserController{}
}

var userServices = service.NewUserService()

func (u UserController) GetInfo(ctx *gin.Context) {
	// 获取用户信息
	claims := jwt.ExtractClaims(ctx)
	id := int(claims["id"].(float64))
	user := model.User{}
	if err := userServices.Retrieve(&user, id); err != nil {
		util.Fail(ctx, gin.H{}, "未找到记录")
		return
	}
	util.Success(ctx, gin.H{"user": user}, "获取用户信息成功")
}

func (u UserController) ChangePassword(ctx *gin.Context) {
	// 用户修改密码
	passwordForm := form.UserChangePasswordForm{}
	var err error
	if err = ctx.ShouldBindJSON(&passwordForm); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "参数错误")
		return
	}
	if passwordForm.Password1 != passwordForm.Password2 {
		util.Fail(ctx, gin.H{}, "两次密码不一致")
		return
	}
	user := model.User{}
	instance := middleware.JWTAuthMiddleware.IdentityHandler(ctx).(*model.User)
	if err = userServices.Retrieve(&user, int(instance.ID)); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "未找到记录")
		return
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwordForm.Password)); err != nil {
		util.Fail(ctx, gin.H{}, "原密码错误")
		return
	}
	var hashedPassword []byte
	if hashedPassword, err = bcrypt.GenerateFromPassword([]byte(passwordForm.Password1), bcrypt.DefaultCost); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "生成密码错误")
		return
	}
	payload := map[string]interface{}{
		"password": string(hashedPassword),
	}
	user.Password = string(hashedPassword)
	if err = userServices.Update(&user, payload); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "修改密码错误")
		return
	}
	util.Success(ctx, gin.H{}, "修改密码成功")
}

func (u UserController) List(ctx *gin.Context) {
	userListQuery := form.UserListQuery{}
	if err := ctx.ShouldBindQuery(&userListQuery); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, nil, "参数错误")
		return
	}
	var users []model.User
	page, size := userListQuery.Page, userListQuery.Size
	var total int64
	query := make(map[string]interface{})
	if userListQuery.Username != "" {
		query["username"] = userListQuery.Username
	}
	if userListQuery.Name != "" {
		query["name"] = userListQuery.Name
	}
	if err := userServices.List(&users, page, size, &total, query); err != nil {
		log.Println(err.Error())
		util.Success(ctx, nil, "获取数据失败")
		return
	}
	util.Success(ctx, gin.H{
		"page":    page,
		"size":    size,
		"results": users,
		"total":   total,
	}, "")
}

func (u UserController) Retrieve(ctx *gin.Context) {
	var id int
	var err error
	if id, err = strconv.Atoi(ctx.Param("id")); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "ID错误")
		return
	}
	user := model.User{}
	if err = userServices.Retrieve(&user, id); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "未找到记录")
		return
	}

	util.Success(ctx, gin.H{"user": user}, "获取用户信息成功")
}

func (u UserController) Create(ctx *gin.Context) {
	createForm := form.UserCreateForm{}
	var err error
	if err = ctx.ShouldBindJSON(&createForm); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "参数错误")
		return
	}
	if createForm.Password1 != createForm.Password2 {
		util.Fail(ctx, gin.H{}, "两次密码不一致")
		return
	}
	var hashedPassword []byte
	if hashedPassword, err = bcrypt.GenerateFromPassword([]byte(createForm.Password1), bcrypt.DefaultCost); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "生成密码错误")
		return
	}
	user := model.User{
		Username:  createForm.Username,
		Nickname:  createForm.Nickname,
		Password:  string(hashedPassword),
		Telephone: createForm.Telephone,
		Email:     createForm.Email,
	}
	if err = userServices.Create(&user); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "创建用户失败")
		return
	}
	// 更新为公共角色
	//if err = userServices.UpdateRole(&user, []int{1}); err != nil {
	//	log.Println(err.Error())
	//	util.Fail(ctx, gin.H{}, "更新角色失败")
	//	return
	//}
	util.Success(ctx, gin.H{}, "创建用户成功")
}

func (u UserController) Update(ctx *gin.Context) {
	var id int
	var err error
	id, err = strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "ID错误")
		return
	}
	updateForm := form.UserUpdateForm{}
	if err = ctx.ShouldBindJSON(&updateForm); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "参数错误")
		return
	}
	user := model.User{}
	if err = userServices.Retrieve(&user, id); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "record not found")
		return
	}
	payload := map[string]interface{}{
		"nickname":  updateForm.Nickname,
		"telephone": updateForm.Telephone,
		"email":     updateForm.Email,
		//"roles":     updateForm.Roles,
	}
	if err = userServices.Update(&user, payload); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "更新用户失败")
		return
	}
	if err = userServices.UpdateRole(&user, updateForm.Roles); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "更新用户失败")
		return
	}
	util.Success(ctx, gin.H{}, "更新成功")
}

func (u UserController) Destroy(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "param error")
		return
	}
	if err = userServices.Destroy(id); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "delete fail")
		return
	}
	util.Success(ctx, gin.H{}, "delete success")
}
