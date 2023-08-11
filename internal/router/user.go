package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zyuanx/research-sys/internal/controller"
	"github.com/zyuanx/research-sys/internal/middleware"
)

func RegisterUserRouter(r *gin.RouterGroup) {

	userController := controller.NewUserController()
	r.POST("/login", middleware.JWTAuthMiddleware.LoginHandler)

	group := r.Group("")
	group.Use(middleware.AuthToken())
	{
		group.POST("/logout", middleware.JWTAuthMiddleware.LogoutHandler)
		group.GET("/refresh_token", middleware.JWTAuthMiddleware.RefreshHandler)
		group.GET("/info", userController.GetInfo)

		group.PUT("/change/password", userController.ChangePassword)

		group.GET("", userController.List)
		group.GET("/:id", userController.Retrieve)
		group.POST("", userController.Create)
		group.PUT("/:id", userController.Update)
		group.DELETE("/:id", userController.Destroy)
	}
}
