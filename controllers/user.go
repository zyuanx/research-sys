package controllers

import (
	"gin-research-sys/middlewares"
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
	GetInfo(ctx *gin.Context)

	List(ctx *gin.Context)
	Retrieve(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Destroy(ctx *gin.Context)
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

func (u UserController) GetInfo(ctx *gin.Context) {
	user := middlewares.JWTAuthMiddleware.IdentityHandler(ctx).(models.User)
	err := userServices.UserInfo(&user)
	if err != nil {
		return
	}
	res.Success(ctx, gin.H{"user": user}, "")
}

func (u UserController) List(ctx *gin.Context) {
	pg := req.PaginationQuery{}
	if err := ctx.ShouldBindQuery(&pg); err != nil {
		res.Success(ctx, nil, err.Error())
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
		res.Fail(ctx, gin.H{}, err.Error())
	}
	user := models.User{}
	if err = userServices.Retrieve(&user, id); err != nil {
		res.Fail(ctx, gin.H{}, err.Error())
		return
	}
	res.Success(ctx, gin.H{"user": user}, "")
}

func (u UserController) Create(ctx *gin.Context) {
	panic("implement me")
}

func (u UserController) Update(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		res.Fail(ctx, gin.H{}, err.Error())
		return
	}
	roleUpdateValidate := req.RoleUpdateValidate{}
	if err = ctx.ShouldBindJSON(&roleUpdateValidate); err != nil {
		res.Fail(ctx, gin.H{}, err.Error())
		return
	}
	role := models.Role{}
	if err = roleServices.Retrieve(&role, id); err != nil {
		res.Fail(ctx, gin.H{}, err.Error())
		return
	}
	if err = roleServices.Update(&role, &roleUpdateValidate); err != nil {
		res.Fail(ctx, gin.H{}, err.Error())
		return
	}
}

func (u UserController) Destroy(ctx *gin.Context) {
	panic("implement me")
}
