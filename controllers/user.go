package controllers

import (
	"gin-research-sys/controllers/req"
	"gin-research-sys/controllers/res"
	"gin-research-sys/middlewares"
	"gin-research-sys/models"
	"gin-research-sys/services"
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

var userServices = services.NewUserService()

func (u UserController) GetInfo(ctx *gin.Context) {
	id := middlewares.JWTAuthMiddleware.IdentityHandler(ctx).(models.User).ID
	user := models.User{}
	if err := userServices.Retrieve(&user, int(id)); err != nil {
		log.Println(err.Error())
		res.Fail(ctx, gin.H{}, "get info error")
	}
	var roles []string
	if err := userServices.ListRole(&user, &roles); err != nil {
		log.Println(err.Error())
		res.Fail(ctx, gin.H{}, "get roles error")
	}

	res.Success(ctx, gin.H{"user": gin.H{
		"username":  user.Username,
		"nickname":  user.Nickname,
		"telephone": user.Telephone,
		"email":     user.Email,
		"roles":     roles,
	}}, "")
}

func (u UserController) ResetPassword(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		res.Fail(ctx, gin.H{}, "param error")
	}
	user := models.User{}
	if err = userServices.Retrieve(&user, id); err != nil {
		res.Fail(ctx, gin.H{}, "record not found")
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err.Error())
		res.Success(ctx, gin.H{}, "generate error")
		return
	}
	user.Password = string(hashedPassword)
	if err = userServices.Update(&user); err != nil {
		log.Println(err.Error())
		res.Success(ctx, gin.H{}, "update fail")
		return
	}
	res.Success(ctx, gin.H{}, "update success")
}

func (u UserController) ChangePassword(ctx *gin.Context) {

	passwordReq := req.UserChangePasswordReq{}
	if err := ctx.ShouldBindJSON(&passwordReq); err != nil {
		log.Println(err.Error())
		res.Fail(ctx, gin.H{}, "payload is error")
		return
	}
	if passwordReq.Password1 != passwordReq.Password2 {
		res.Fail(ctx, gin.H{}, "the two passwords don't match")
		return
	}
	user := models.User{}
	ins := middlewares.JWTAuthMiddleware.IdentityHandler(ctx).(models.User)
	if err := userServices.Retrieve(&user, int(ins.ID)); err != nil {
		log.Println(err.Error())
		res.Fail(ctx, gin.H{}, "record not found")
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwordReq.Password)); err != nil {
		res.Fail(ctx, gin.H{}, "the password is wrong")
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passwordReq.Password1), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err.Error())
		res.Success(ctx, gin.H{}, "generate error")
		return
	}
	user.Password = string(hashedPassword)
	if err := userServices.Update(&user); err != nil {
		log.Println(err.Error())
		res.Success(ctx, gin.H{}, "update fail")
		return
	}
	res.Success(ctx, gin.H{}, "update success")

}

func (u UserController) List(ctx *gin.Context) {
	pg := req.PaginationQuery{}
	if err := ctx.ShouldBindQuery(&pg); err != nil {
		res.Fail(ctx, nil, err.Error())
		return
	}
	var users []models.User
	var total int64
	if err := userServices.List(pg.Page, pg.Size, &users, &total); err != nil {
		res.Success(ctx, nil, err.Error())
		return
	}
	res.Success(ctx, gin.H{
		"page":    pg.Page,
		"size":    pg.Size,
		"results": users,
		"total":   total,
	}, "")
}

func (u UserController) Retrieve(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		res.Fail(ctx, gin.H{}, "param error")
		return
	}
	user := models.User{}
	if err = userServices.Retrieve(&user, id); err != nil {
		res.Fail(ctx, gin.H{}, "record not found")
		return
	}
	var roles []int
	if err = userServices.ListRole2(&user, &roles); err != nil {
		res.Fail(ctx, gin.H{}, "get roles error")
	}

	res.Success(ctx, gin.H{"user": gin.H{
		"id":        user.ID,
		"username":  user.Username,
		"nickname":  user.Nickname,
		"telephone": user.Telephone,
		"email":     user.Email,
		"roles":     roles,
	}}, "")
}

func (u UserController) Create(ctx *gin.Context) {
	ucq := req.UserCreateReq{}

	if err := ctx.ShouldBindJSON(&ucq); err != nil {
		log.Println(err.Error())
		res.Fail(ctx, gin.H{}, "payload is error")
		return
	}
	if ucq.Password1 != ucq.Password2 {
		res.Fail(ctx, gin.H{}, "the two passwords don't match")
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(ucq.Password1), bcrypt.DefaultCost)
	if err != nil {
		res.Fail(ctx, gin.H{}, err.Error())
		return
	}
	user := models.User{
		Username:  ucq.Username,
		Nickname:  ucq.Nickname,
		Password:  string(hashedPassword),
		Telephone: ucq.Telephone,
		Email:     ucq.Email,
	}
	if err = userServices.Create(&user); err != nil {
		log.Println(err.Error())
		res.Fail(ctx, gin.H{}, "create fail")
		return
	}
	res.Success(ctx, gin.H{}, "create success")
}

func (u UserController) Update(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Println(err.Error())
		res.Fail(ctx, gin.H{}, "param error")
		return
	}
	updateReq := req.UserUpdateReq{}
	if err = ctx.ShouldBindJSON(&updateReq); err != nil {
		log.Println(err.Error())
		res.Fail(ctx, gin.H{}, "payload error")
		return
	}
	user := models.User{}
	if err = userServices.Retrieve(&user, id); err != nil {
		log.Println(err.Error())
		res.Fail(ctx, gin.H{}, "record not found")
		return
	}
	user.Nickname = updateReq.Nickname
	user.Telephone = updateReq.Telephone
	user.Email = updateReq.Email
	if err = userServices.Update(&user); err != nil {
		log.Println(err.Error())
		res.Fail(ctx, gin.H{}, "update fail")
		return
	}
	if err = userServices.UpdateRole(&user, updateReq.Roles); err != nil {
		log.Println(err.Error())
		res.Fail(ctx, gin.H{}, "update fail")
		return
	}
	res.Success(ctx, gin.H{}, "update success")
}

func (u UserController) Destroy(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Println(err.Error())
		res.Fail(ctx, gin.H{}, "param error")
		return
	}
	if err = userServices.Destroy(id); err != nil {
		log.Println(err.Error())
		res.Fail(ctx, gin.H{}, "delete fail")
		return
	}
	res.Success(ctx, gin.H{}, "delete success")
}
