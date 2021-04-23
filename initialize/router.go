package initialize

import (
	"gin-research-sys/routers"
	"gin-research-sys/routers/api"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	r := gin.Default()
	r = routers.InitRouter()

	apiGroup := r.Group("/api")

	api.RegisterUserRouter(apiGroup.Group("/user"))
	return r
}
