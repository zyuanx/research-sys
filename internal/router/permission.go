package router

import (
	"gin-research-sys/internal/controller"
	"gin-research-sys/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterPermissionRouter(r *gin.RouterGroup) {

	control := controller.NewPermissionController()

	group := r.Group("")
	group.Use(middleware.JWTAuthMiddleware.MiddlewareFunc())
	group.GET("", control.List)
	group.GET("/:id", control.Retrieve)
	group.POST("", control.Create)
	group.PUT("/:id", control.Update)
	group.DELETE("/:id", control.Destroy)

}
