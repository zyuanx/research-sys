package initialize

import (
	_ "gin-research-sys/docs"
	"gin-research-sys/routers"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Routers() *gin.Engine {
	r := gin.Default()

	apiGroup := r.Group("/api")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routers.RegisterUserRouter(apiGroup.Group("/user"))
	routers.RegisterRoleRouter(apiGroup.Group("/role"))
	routers.RegisterResearchRouter(apiGroup.Group("/research"))
	return r
}
