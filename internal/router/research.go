package router

import (
	"gin-research-sys/internal/controller"
	"gin-research-sys/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterResearchRouter(r *gin.RouterGroup) {
	researchController := controller.NewResearchController()

	group := r.Group("")
	group.Use(middleware.JWTAuthMiddleware.MiddlewareFunc())
	group.GET("", researchController.List)
	group.GET("/:id", researchController.Retrieve)
	r.GET("/open/:id", researchController.RetrieveOpen)
	group.POST("", researchController.Create)
	group.PUT("/:id", researchController.Update)
	group.DELETE("/:id", researchController.Destroy)
	group.GET("/square", researchController.Square)
	r.GET("/export/:id", researchController.DownloadExcel)
}
