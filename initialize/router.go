package initialize

import (
	_ "gin-research-sys/docs"
	"gin-research-sys/middlewares"
	"gin-research-sys/pkg/global"
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

	auth := r.Group("/api")
	{
		// 模拟添加一条Policy策略
		auth.POST("acs", func(c *gin.Context) {
			subject := "tom"
			object := "/api/routers"
			action := "POST"
			cacheName := subject + object + action
			result, _ := global.Enforcer.AddPolicy(subject, object, action)
			if result {
				// 清除缓存
				_ = global.Cache.Delete(cacheName)
				c.JSON(200, gin.H{"code": 200, "message": "add success"})
			} else {
				c.JSON(200, gin.H{"code": 200, "message": "add fail"})
			}
		})
		// 模拟删除一条Policy策略
		auth.DELETE("acs/:id", func(c *gin.Context) {
			result, _ := global.Enforcer.RemovePolicy("tom", "/api/routers", "POST")
			if result {
				// 清除缓存 代码省略
				c.JSON(200, gin.H{"code": 200, "message": "delete Policy success"})
			} else {
				c.JSON(200, gin.H{"code": 200, "message": "delete Policy fail"})
			}
		})
		// 获取路由列表
		auth.POST("/routers", middlewares.Privilege(), func(c *gin.Context) {
			type data struct {
				Method string `json:"method"`
				Path   string `json:"path"`
			}
			var datas []data
			rou:= r.Routes()
			for _, v := range rou {
				var temp data
				temp.Method = v.Method
				temp.Path = v.Path
				datas = append(datas, temp)
			}
			c.JSON(200, gin.H{"code": 200, "data": datas})
			return
		})
	}


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
