package api

import (
	"gin-research-sys/controllers"
	"gin-research-sys/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterUserRouter(r *gin.RouterGroup) {

	userController := controllers.NewUserController()

	r.POST("/register", userController.Register)
	r.POST("/login", middlewares.JWTAuthMiddleware.LoginHandler)
	r.POST("/logout", middlewares.JWTAuthMiddleware.LogoutHandler)
	r.GET("/refresh_token", middlewares.JWTAuthMiddleware.RefreshHandler)
	//
	auth := r.Group("")
	auth.Use(middlewares.JWTAuthMiddleware.MiddlewareFunc())
	{
		auth.GET("/info/:id", userController.GetInfo)
	}
}
