package router

import (
	"gin-research-sys/internal/controller"
	"gin-research-sys/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRecordRouter(r *gin.RouterGroup) {

	recordController := controller.NewRecordController()

	group := r.Group("")
	group.Use(middleware.JWTAuthMiddleware.MiddlewareFunc())
	group.GET("", recordController.List)
	group.GET("/:id", recordController.Retrieve)
	group.POST("", recordController.Create)
}
