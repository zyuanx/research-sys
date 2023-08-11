package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zyuanx/research-sys/internal/controller"
	"github.com/zyuanx/research-sys/internal/middleware"
)

func RegisterRecordRouter(r *gin.RouterGroup) {

	recordController := controller.NewRecordController()
	group := r.Group("")
	group.Use(middleware.AuthToken())
	group.GET("", recordController.List)
	group.GET("/:id", recordController.Retrieve)
	group.POST("", recordController.Create)
	group.GET("/filled/:id", recordController.Filled)

	openRecordController := controller.NewOpenRecordController()
	group.GET("/open", openRecordController.List)
	group.GET("/open/:id", openRecordController.Retrieve)
	r.POST("/open", openRecordController.Create)
}
