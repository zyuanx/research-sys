package controllers

import (
	"gin-research-sys/models"
	"gin-research-sys/pkg/req"
	"gin-research-sys/pkg/res"
	"gin-research-sys/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

type UserController struct {
}

func NewUserController() IUserController {
	return UserController{}
}

type IUserController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	GetInfo(ctx *gin.Context)
}

var userServices = services.NewUserService()

func (u UserController) Register(ctx *gin.Context) {
	registerValidator := req.RegisterValidator{}
	if err := ctx.ShouldBindJSON(&registerValidator); err != nil {
		res.Fail(ctx, gin.H{}, err.Error())
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerValidator.Password), bcrypt.DefaultCost)
	if err != nil {
		res.Fail(ctx, gin.H{}, err.Error())
		return
	}
	user := models.User{
		Username: registerValidator.Username,
		Password: string(hashedPassword),
	}

	err = userServices.UserRegister(&user)
	if err != nil {
		res.Fail(ctx, gin.H{}, err.Error())
		return
	}
	res.Success(ctx, gin.H{"user": user}, "")
}

func (u UserController) Login(ctx *gin.Context) {
	login := req.LoginValidator{}
	if err := ctx.ShouldBindJSON(&login); err != nil {
		res.Fail(ctx, gin.H{}, err.Error())
		return
	}
	user := models.User{
		Username: login.Username,
		Password: login.Password,
	}
	if err := userServices.UserLogin(&user); err != nil {
		res.Fail(ctx, gin.H{}, err.Error())
		return
	}
	password := user.Password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		res.Fail(ctx, gin.H{}, err.Error())
		return
	}
	res.Success(ctx, gin.H{"user": user}, "")
}

func (u UserController) GetInfo(ctx *gin.Context) {
	//user := middlewares.JWTAuthMiddleware.IdentityHandler(ctx)
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		res.Fail(ctx, gin.H{}, err.Error())
		return
	}
	user := models.User{}
	if err := userServices.UserInfo(&user, id); err != nil {
		res.Fail(ctx, gin.H{}, err.Error())
		return
	}
	res.Success(ctx, gin.H{
		"user": res.InfoSerializer(user),
	}, "")
}
