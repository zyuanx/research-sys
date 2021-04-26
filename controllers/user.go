package controllers

import (
	"gin-research-sys/controllers/request"
	"gin-research-sys/controllers/response"
	"gin-research-sys/models"
	"gin-research-sys/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

type IUserController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	GetInfo(ctx *gin.Context)
}
type UserController struct {
}

func NewUserController() IUserController {
	return UserController{}
}
func (u UserController) Register(ctx *gin.Context) {
	registerValidator := request.RegisterValidator{}
	if err := ctx.ShouldBindJSON(&registerValidator); err != nil {
		response.Fail(ctx, gin.H{}, err.Error())
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerValidator.Password), bcrypt.DefaultCost)
	if err != nil {
		response.Fail(ctx, gin.H{}, err.Error())
		return
	}
	user := models.User{
		Username: registerValidator.Username,
		Password: string(hashedPassword),
	}

	err = services.UserRegister(&user)
	if err != nil {
		response.Fail(ctx, gin.H{}, err.Error())
		return
	}
	response.Success(ctx, gin.H{"user": user}, "")
}

func (u UserController) Login(ctx *gin.Context) {
	login := request.LoginValidator{}
	if err := ctx.ShouldBindJSON(&login); err != nil {
		response.Fail(ctx, gin.H{}, err.Error())
		return
	}
	user := models.User{
		Username: login.Username,
		Password: login.Password,
	}
	if err := services.UserLogin(&user); err != nil {
		response.Fail(ctx, gin.H{}, err.Error())
		return
	}
	password := user.Password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		response.Fail(ctx, gin.H{}, err.Error())
		return
	}
	response.Success(ctx, gin.H{"user": user}, "")
}

func (u UserController) GetInfo(ctx *gin.Context) {
	//user := middlewares.JWTAuthMiddleware.IdentityHandler(ctx)
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		response.Fail(ctx, gin.H{}, err.Error())
		return
	}
	user := models.User{}
	if err := services.UserInfo(&user, id); err != nil {
		response.Fail(ctx, gin.H{}, err.Error())
		return
	}
	response.Success(ctx, gin.H{
		"user": response.InfoSerializer(user),
	}, "")
}
