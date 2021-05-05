package routers

import (
	"gin-research-sys/controllers"
	"gin-research-sys/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterUserRouter(r *gin.RouterGroup) {

	userController := controllers.NewUserController()
	r.POST("/login", middlewares.JWTAuthMiddleware.LoginHandler)

	//
	auth := r.Group("")
	auth.Use(middlewares.JWTAuthMiddleware.MiddlewareFunc())
	{
		auth.POST("/logout", middlewares.JWTAuthMiddleware.LogoutHandler)
		auth.GET("/refresh_token", middlewares.JWTAuthMiddleware.RefreshHandler)
		auth.GET("/info", userController.GetInfo)
		auth.GET("", userController.List)
		auth.GET("/:id", userController.Retrieve)
		auth.POST("", userController.Create)
		auth.PUT("/:id", userController.Update)
		auth.DELETE("/:id", userController.Destroy)
	}
}
