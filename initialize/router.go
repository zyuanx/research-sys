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
	//r := gin.New()
	//
	//r.Use(ginzap.Ginzap(zap.L(), time.RFC3339, true))
	//r.Use(ginzap.RecoveryWithZap(zap.L(), true))

	r.Use(middlewares.Cors())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	apiGroup := r.Group("/api")
	routers.RegisterUserRouter(apiGroup.Group("/user"))
	routers.RegisterRoleRouter(apiGroup.Group("/role"))
	routers.RegisterPermissionRouter(apiGroup.Group("/permission"))
	routers.RegisterResearchRouter(apiGroup.Group("/research"))
	routers.RegisterRecordRouter(apiGroup.Group("/record"))

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
