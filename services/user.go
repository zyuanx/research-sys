package services

import (
	"gin-research-sys/common"
	"gin-research-sys/models"
	"gin-research-sys/services/request"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func Register(ctx *gin.Context) {
	registerValidator := request.RegisterValidator{}
	if err := ctx.ShouldBindJSON(&registerValidator); err != nil {
		common.Fail(ctx, 400, gin.H{}, err.Error())
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerValidator.Password), bcrypt.DefaultCost)
	if err != nil {
		common.Response(ctx,
			http.StatusInternalServerError,
			500,
			nil,
			"加密错误")
		return
	}
	user := models.User{
		Username: registerValidator.Username,
		Password: string(hashedPassword),
	}
	common.DB.Create(&user)
	common.Success(ctx, 200, gin.H{"user": user}, "success")
}

//func userLogin(ctx *gin.Context) {
//	login := LoginValidator{}
//	if err := ctx.ShouldBindJSON(&login); err != nil {
//		response.Fail(ctx, 400, gin.H{}, err.Error())
//		return
//	}
//	user := &User{
//		Username: login.Username,
//		Password: login.Password,
//	}
//	password := user.Password
//	result := db.Where("username = ?", user.Username).First(&user)
//	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
//
//}

func UserAdd(ctx *gin.Context) {
	u := models.User{}
	if err := ctx.ShouldBindJSON(&u); err != nil {
		common.Fail(ctx, 400, gin.H{}, err.Error())
		return
	}

	if err := u.Add(); err != nil {
		common.Fail(ctx, 400, gin.H{}, err.Error())
	} else {
		common.Success(ctx, 200, gin.H{"user": u}, "user add success")
	}
}

func UserInfo(ctx *gin.Context) {
	claims := jwt.ExtractClaims(ctx)
	log.Println(claims)
	//u, _ := ctx.Get("user")
	//b, err := json.Marshal(u)
	//if err != nil {
	//	common.Fail(ctx, 400, gin.H{}, err.Error())
	//} else {
	//	common.Success(ctx, 200, gin.H{"user": b}, "success")
	//}
}

func UserTest(ctx *gin.Context)  {
	common.Success(ctx, 200, gin.H{}, "UserTest")
}