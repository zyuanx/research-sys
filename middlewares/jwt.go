package middlewares

import (
	"gin-research-sys/models"
	"gin-research-sys/services"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

var identityKey = "id"
var JWTAuthMiddleware *jwt.GinJWTMiddleware

func init() {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour * 24,
		MaxRefresh:  time.Hour * 24,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					identityKey: v.ID,
					"username":  v.Username,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &models.User{
				Model: gorm.Model{
					ID: uint(claims[identityKey].(float64)),
				},
				Username: claims["username"].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			user, err := services.UserLogin(c)
			if err != nil {
				return nil, jwt.ErrFailedAuthentication
			}
			return user, nil

		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if _, ok := data.(*models.User); ok {
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
