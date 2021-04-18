package user

import (
	"encoding/json"
	"gin-research-sys/common/response"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Register(ctx *gin.Context) {
	registerValidator := RegisterValidator{}
	if err := ctx.ShouldBindJSON(&registerValidator); err != nil {
		response.Fail(ctx, 400, gin.H{}, err.Error())
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerValidator.Password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx,
			http.StatusInternalServerError,
			500,
			nil,
			"加密错误")
		return
	}
	user := User{
		Username: registerValidator.Username,
		Password: string(hashedPassword),
	}
	db.Create(&user)
	response.Success(ctx, 200, gin.H{"user": user}, "success")
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

func userAdd(ctx *gin.Context) {
	u := User{}
	if err := ctx.ShouldBindJSON(&u); err != nil {
		response.Fail(ctx, 400, gin.H{}, err.Error())
		return
	}

	if err := u.add(); err != nil {
		response.Fail(ctx, 400, gin.H{}, err.Error())
	} else {
		response.Success(ctx, 200, gin.H{"user": u}, "user add success")
	}
}

func userInfo(ctx *gin.Context) {
	u, _ := ctx.Get("user")
	b, err := json.Marshal(u)
	if err != nil {
		response.Fail(ctx, 400, gin.H{}, err.Error())
	} else {
		response.Success(ctx, 200, gin.H{"user": b}, "success")
	}
}
