package initialize

import (
	_ "gin-research-sys/docs"
	"gin-research-sys/internal/middleware"
	"gin-research-sys/internal/router"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Router() *gin.Engine {
	r := gin.Default()
	//r := gin.New()
	//
	//r.Use(ginzap.Ginzap(zap.L(), time.RFC3339, true))
	//r.Use(ginzap.RecoveryWithZap(zap.L(), true))

	r.Use(middleware.Cors())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	apiGroup := r.Group("/api")
	router.RegisterUserRouter(apiGroup.Group("/user"))
	router.RegisterRoleRouter(apiGroup.Group("/role"))
	router.RegisterPermissionRouter(apiGroup.Group("/permission"))
	router.RegisterResearchRouter(apiGroup.Group("/research"))
	router.RegisterRecordRouter(apiGroup.Group("/record"))

	//user := r.Group("/api/v1")
	//// use cabin middleware
	//user.Use(middleware.Privilege())
	//{
	//	user.POST("user", func(c *gin.Context) {
	//		c.JSON(200, gin.H{"code": 200, "message": "user add success"})
	//	})
	//	user.DELETE("user/:id", func(c *gin.Context) {
	//		id := c.Param("id")
	//		c.JSON(200, gin.H{"code": 200, "message": "user delete success " + id})
	//	})
	//	user.PUT("user/:id", func(c *gin.Context) {
	//		id := c.Param("id")
	//		c.JSON(200, gin.H{"code": 200, "message": "user update success " + id})
	//	})
	//	user.GET("user/:id", func(c *gin.Context) {
	//		id := c.Param("id")
	//		c.JSON(200, gin.H{"code": 200, "message": "user Get success " + id})
	//	})
	//}

	return r
}
