package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zyuanx/research-sys/internal/middleware"
)

func RegisterUserRouter(r *gin.RouterGroup) {
	group := r.Group("user")
	group.POST("/login", c.UserLogin)

	group.Use(middleware.AuthToken())
	{
		// group.POST("/logout", middleware.JWTAuthMiddleware.LogoutHandler)
		// group.GET("/refresh_token", middleware.JWTAuthMiddleware.RefreshHandler)
		group.GET("/info", c.UserGetInfo)

		group.PUT("/change/password", c.UserChangePassword)

		group.GET("", c.UserList)
		group.GET("/:id", c.UserRetrieve)
		group.POST("", c.UserCreate)
		group.PUT("/:id", c.UserUpdate)
		group.DELETE("/:id", c.UserDelete)
	}
}
