package middleware

import (
	"errors"
	"gin-research-sys/internal/model"
	"gin-research-sys/internal/service"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

var identityKey = "id"
var JWTAuthMiddleware *jwt.GinJWTMiddleware

func init() {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "research",
		Key:         []byte("QhYTOVSfGa0xFE4sctH6lj7UuZRiq5m2"),
		Timeout:     time.Hour * 24,
		MaxRefresh:  time.Hour * 24,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				return jwt.MapClaims{
					"username":  v.Username,
					identityKey: v.ID,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			// 此处返回值类型 interface{} 与 payloadFunc 和 Authenticator 的data类型必须一致,
			// 否则会导致授权失败
			return &model.User{
				Username: claims["username"].(string),
				BaseModel: model.BaseModel{
					ID: uint(claims[identityKey].(float64)),
				},
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			login := LoginReq{}
			if err := c.ShouldBindJSON(&login); err != nil {
				return nil, errors.New("payload is error")
			}
			user := model.User{
				Username: login.Username,
				Password: login.Password,
			}
			userService := service.NewUserService()
			if err := userService.FindUserByUsername(&user); err != nil {
				return nil, err
			}
			err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
			if err != nil {
				return nil, jwt.ErrFailedAuthentication
			}
			return &user, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if _, ok := data.(*model.User); ok {
				return true
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"data": gin.H{
					"token":  token,
					"expire": expire.Format(time.RFC3339),
				},
			})
		},
		RefreshResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"data": gin.H{
					"token":  token,
					"expire": expire.Format(time.RFC3339),
				},
			})
		},
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	JWTAuthMiddleware = authMiddleware
}
