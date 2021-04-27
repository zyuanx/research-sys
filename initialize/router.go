package initialize

import (
	_ "gin-research-sys/docs"
	"gin-research-sys/routers/api"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Routers() *gin.Engine {
	r := gin.Default()

	apiGroup := r.Group("/api")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api.RegisterUserRouter(apiGroup.Group("/user"))
	api.RegisterRoleRouter(apiGroup.Group("/role"))
	return r
}
