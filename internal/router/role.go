package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zyuanx/research-sys/internal/controller"
	"github.com/zyuanx/research-sys/internal/middleware"
)

func RegisterRoleRouter(r *gin.RouterGroup) {

	roleController := controller.NewRoleController()

	group := r.Group("")
	group.Use(middleware.AuthToken())
	group.GET("", roleController.List)
	group.GET("/:id", roleController.Retrieve)
	group.POST("", roleController.Create)
	group.PUT("/:id", roleController.Update)
	group.DELETE("/:id", roleController.Destroy)

}
