package initialize

import (
	"gin-research-sys/routers/api"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	r := gin.Default()

	apiGroup := r.Group("/api")

	api.RegisterUserRouter(apiGroup.Group("/user"))
	api.RegisterRoleRouter(apiGroup.Group("/role"))
	return r
}
