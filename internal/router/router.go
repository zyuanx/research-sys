package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zyuanx/research-sys/internal/controller"
	"github.com/zyuanx/research-sys/internal/middleware"
	"github.com/zyuanx/research-sys/internal/service"
)

func SetupRouter(r *gin.Engine, s *service.Service) *gin.Engine {

	r.Use(middleware.CORS())
	// r.Use(middleware.RequestId())
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	controller := controller.NewController(s)
	apiGroup := r.Group("/api")
	RegisterUserRouter(controller, apiGroup.Group("/user"))
	// RegisterRoleRouter(controller, apiGroup.Group("/role"))
	// RegisterPermissionRouter(controller, apiGroup.Group("/permission"))
	// RegisterResearchRouter(controller, controllerapiGroup.Group("/research"))
	// RegisterRecordRouter(controller, apiGroup.Group("/record"))
	return r
}
