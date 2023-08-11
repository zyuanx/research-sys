package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zyuanx/research-sys/internal/controller"
	"github.com/zyuanx/research-sys/internal/middleware"
)

func RegisterPermissionRouter(r *gin.RouterGroup) {

	control := controller.NewPermissionController()

	group := r.Group("")
	group.Use(middleware.AuthToken())
	group.GET("", control.List)
	group.GET("/:id", control.Retrieve)
	group.POST("", control.Create)
	group.PUT("/:id", control.Update)
	group.DELETE("/:id", control.Destroy)

}
