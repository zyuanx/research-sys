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
	ResetPassword(ctx *gin.Context)
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

func (u UserController) GetInfo(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	id := int(claims["id"].(float64))
	user := model.User{}
	if err := userServices.Retrieve(&user, id); err != nil {
		util.Fail(c, gin.H{}, "record not found")
		return
	}
	var roles []string
	for _, value := range user.Roles {
		roles = append(roles, value.Title)
	}
	util.Success(c, gin.H{"user": gin.H{
		"username":  user.Username,
		"nickname":  user.Nickname,
		"telephone": user.Telephone,
		"email":     user.Email,
		"roles":     roles,
	}}, "")
}

func (u UserController) ResetPassword(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Println(err.Error())
		util.Fail(c, gin.H{}, "param error")
		return
	}
	user := model.User{}
	if err = userServices.Retrieve(&user, id); err != nil {
		log.Println(err.Error())
		util.Fail(c, gin.H{}, "record not found")
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err.Error())
		util.Fail(c, gin.H{}, "generate password error")
		return
	}
	user.Password = string(hashedPassword)
	if err = userServices.Update(&user); err != nil {
		log.Println(err.Error())
		util.Fail(c, gin.H{}, "reset password fail")
		return
	}
	util.Success(c, gin.H{}, "reset password success")
}

func (u UserController) ChangePassword(ctx *gin.Context) {
	passwordForm := form.UserChangePasswordForm{}
	if err := ctx.ShouldBindJSON(&passwordForm); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "payload error")
		return
	}
	if passwordForm.Password1 != passwordForm.Password2 {
		util.Fail(ctx, gin.H{}, "the two passwords don't match")
		return
	}
	user := model.User{}
	ins := middleware.JWTAuthMiddleware.IdentityHandler(ctx).(model.User)
	if err := userServices.Retrieve(&user, int(ins.ID)); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "record not found")
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwordForm.Password)); err != nil {
		util.Fail(ctx, gin.H{}, "the password is wrong")
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passwordForm.Password1), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "password generate error")
		return
	}
	user.Password = string(hashedPassword)
	if err := userServices.Update(&user); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "update password fail")
		return
	}
	util.Success(ctx, gin.H{}, "update password success")

}

func (u UserController) List(ctx *gin.Context) {
	pagination := form.Pagination{}
	if err := ctx.ShouldBindQuery(&pagination); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, nil, "query error")
		return
	}
	var users []model.User
	var total int64
	if err := userServices.List(pagination.Page, pagination.Size, &users, &total); err != nil {
		log.Println(err.Error())
		util.Success(ctx, nil, "list user error")
		return
	}
	util.Success(ctx, gin.H{
		"page":    pagination.Page,
		"size":    pagination.Size,
		"results": users,
		"total":   total,
	}, "")
}

func (u UserController) Retrieve(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "param error")
		return
	}
	user := model.User{}
	if err = userServices.Retrieve(&user, id); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "record not found")
		return
	}
	var roles []int
	if err = userServices.ListRole2(&user, &roles); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "get roles error")
		return
	}

	util.Success(ctx, gin.H{"user": gin.H{
		"id":        user.ID,
		"username":  user.Username,
		"nickname":  user.Nickname,
		"telephone": user.Telephone,
		"email":     user.Email,
		"roles":     roles,
	}}, "")
}

func (u UserController) Create(ctx *gin.Context) {
	createForm := form.UserCreateForm{}

	if err := ctx.ShouldBindJSON(&createForm); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "payload is error")
		return
	}
	if createForm.Password1 != createForm.Password2 {
		util.Fail(ctx, gin.H{}, "the two passwords don't match")
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(createForm.Password1), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "password generate error")
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
		util.Fail(ctx, gin.H{}, "create fail")
		return
	}
	// 更新为公共角色
	if err = userServices.UpdateRole(&user, []int{1}); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "update fail")
		return
	}
	util.Success(ctx, gin.H{}, "create success")
}

func (u UserController) Update(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "param error")
		return
	}
	updateForm := form.UserUpdateForm{}
	if err = ctx.ShouldBindJSON(&updateForm); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "payload error")
		return
	}
	user := model.User{}
	if err = userServices.Retrieve(&user, id); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "record not found")
		return
	}
	user.Nickname = updateForm.Nickname
	user.Telephone = updateForm.Telephone
	user.Email = updateForm.Email
	if err = userServices.Update(&user); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "update fail")
		return
	}
	if err = userServices.UpdateRole(&user, updateForm.Roles); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "update fail")
		return
	}
	util.Success(ctx, gin.H{}, "update success")
}

func (u UserController) Destroy(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
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
