package routers

import (
	api2 "gin-research-sys/routers/api"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {

	router := gin.Default()

	setUpRouter(router)

	return router
}

// RegisterRouter 注册路由
func setUpRouter(router *gin.Engine) {

	api := router.Group("/api")
	{
		api2.RegisterRouter(api.Group("/user"))
	}

}
