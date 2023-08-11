package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zyuanx/research-sys/internal/middleware"
)

func SetupRouter(r *gin.Engine) *gin.Engine {

	r.Use(middleware.CORS())
	// r.Use(middleware.RequestId())
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	apiGroup := r.Group("/api")
	RegisterUserRouter(apiGroup.Group("/user"))
	RegisterRoleRouter(apiGroup.Group("/role"))
	RegisterPermissionRouter(apiGroup.Group("/permission"))
	RegisterResearchRouter(apiGroup.Group("/research"))
	RegisterRecordRouter(apiGroup.Group("/record"))
	return r
}
