package user

import (
	"github.com/gin-gonic/gin"
)

// RegisterRouter 注册路由
func RegisterRouter(r *gin.RouterGroup) {

	r.POST("/register", Register)
	r.POST("/login", JWTAuthMiddleware.LoginHandler)
	r.GET("/refresh_token", JWTAuthMiddleware.RefreshHandler)
	//
	auth := r.Group("")
	auth.Use(JWTAuthMiddleware.MiddlewareFunc())
	{
		auth.POST("/add", userAdd)
		auth.GET("/info/:id", userInfo)

	}
}
