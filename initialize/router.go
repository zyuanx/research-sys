package initialize

import (
	_ "gin-research-sys/docs"
	"gin-research-sys/middlewares"
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

	user := r.Group("/api/v1")
	// 使用访问控制中间件
	user.Use(middlewares.Privilege())
	{
		user.POST("user", func(c *gin.Context) {
			c.JSON(200, gin.H{"code": 200, "message": "user add success"})
		})
		user.DELETE("user/:id", func(c *gin.Context) {
			id := c.Param("id")
			c.JSON(200, gin.H{"code": 200, "message": "user delete success " + id})
		})
		user.PUT("user/:id", func(c *gin.Context) {
			id := c.Param("id")
			c.JSON(200, gin.H{"code": 200, "message": "user update success " + id})
		})
		user.GET("user/:id", func(c *gin.Context) {
			id := c.Param("id")
			c.JSON(200, gin.H{"code": 200, "message": "user Get success " + id})
		})
	}

	return r
}
