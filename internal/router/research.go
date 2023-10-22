package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zyuanx/research-sys/internal/middleware"
)

func RegisterResearchRouter(r *gin.RouterGroup) {

	group := r.Group("research")
	group.Use(middleware.AuthToken())
	{
		group.GET("", c.ResearchList)
		group.GET("/:id", c.ResearchRetrieve)
		group.POST("", c.ResearchCreate)
		group.PUT("/:id", c.ResearchUpdate)
		group.DELETE("/:id", c.ResearchDelete)
		// group.GET("/square", researchController.Square)
		// r.GET("/export/:id", researchController.DownloadExcel)
	}
}
