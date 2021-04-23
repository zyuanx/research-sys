package api

import (
	"gin-research-sys/middlewares"
	"gin-research-sys/services"
	"github.com/gin-gonic/gin"
)

func RegisterUserRouter(r *gin.RouterGroup) {

	r.POST("/register", services.Register)
	r.POST("/login", middlewares.JWTAuthMiddleware.LoginHandler)
	r.POST("/logout", middlewares.JWTAuthMiddleware.LogoutHandler)
	r.GET("/refresh_token", middlewares.JWTAuthMiddleware.RefreshHandler)
	//
	auth := r.Group("")
	auth.Use(middlewares.JWTAuthMiddleware.MiddlewareFunc())
	{
		auth.GET("/info/:id", services.UserInfo)
		auth.GET("/test", services.UserTest)

	}
}
