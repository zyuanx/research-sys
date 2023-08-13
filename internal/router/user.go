package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zyuanx/research-sys/internal/controller"
	"github.com/zyuanx/research-sys/internal/middleware"
)

func RegisterUserRouter(controller controller.Controller, r *gin.RouterGroup) {
	r.POST("/login", controller.UserLogin)

	group := r.Group("")
	group.Use(middleware.AuthToken())
	{
		// group.POST("/logout", middleware.JWTAuthMiddleware.LogoutHandler)
		// group.GET("/refresh_token", middleware.JWTAuthMiddleware.RefreshHandler)
		group.GET("/info", controller.UserGetInfo)

		group.PUT("/change/password", controller.UserChangePassword)

		group.GET("", controller.UserList)
		group.GET("/:id", controller.UserRetrieve)
		group.POST("", controller.UserCreate)
		group.PUT("/:id", controller.UserUpdate)
		group.DELETE("/:id", controller.UserDelete)
	}
}
