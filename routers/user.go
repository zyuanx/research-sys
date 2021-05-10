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
	group := r.Group("")
	group.Use(middlewares.JWTAuthMiddleware.MiddlewareFunc())
	{
		group.POST("/logout", middlewares.JWTAuthMiddleware.LogoutHandler)
		group.GET("/refresh_token", middlewares.JWTAuthMiddleware.RefreshHandler)
		group.GET("/info", userController.GetInfo)

		group.PUT("/reset/password/:id", userController.ResetPassword)
		group.PUT("/change/password", userController.ChangePassword)

		group.GET("", userController.List)
		group.GET("/:id", userController.Retrieve)
		group.POST("", userController.Create)
		group.PUT("/:id", userController.Update)
		group.DELETE("/:id", userController.Destroy)
	}
}
