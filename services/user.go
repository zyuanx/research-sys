package services

import (
	"gin-research-sys/common"
	"gin-research-sys/models"
	"gin-research-sys/pkg/global"
	"gin-research-sys/services/request"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func Register(ctx *gin.Context) (*models.User, error) {
	registerValidator := request.RegisterValidator{}
	if err := ctx.ShouldBindJSON(&registerValidator); err != nil {
		return nil, err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerValidator.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := models.User{
		Username: registerValidator.Username,
		Password: string(hashedPassword),
	}
	global.Mysql.Create(user)
	return &user, nil
}

func UserLogin(ctx *gin.Context) (*models.User, error) {
	login := request.LoginValidator{}
	if err := ctx.ShouldBindJSON(login); err != nil {
		return nil, err
	}
	user := models.User{
		Username: login.Username,
		Password: login.Password,
	}
	password := user.Password
	result := global.Mysql.Where("username = ?", user.Username).First(user)
	if result.Error != nil {
		return nil, result.Error
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}
	return &user, nil

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

func UserTest(ctx *gin.Context) {
	common.Success(ctx, 200, gin.H{}, "UserTest")
}
