package router

import (
	"gin-research-sys/internal/controller"
	"gin-research-sys/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoleRouter(r *gin.RouterGroup) {

	roleController := controller.NewRoleController()

	group := r.Group("")
	group.Use(middleware.JWTAuthMiddleware.MiddlewareFunc())
	group.GET("", roleController.List)
	group.GET("/:id", roleController.Retrieve)
	group.POST("", roleController.Create)
	group.PUT("/:id", roleController.Update)
	group.DELETE("/:id", roleController.Destroy)

}
