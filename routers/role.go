package routers

import (
	"gin-research-sys/controllers"
	"gin-research-sys/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoleRouter(r *gin.RouterGroup) {

	controller := controllers.NewRoleController()

	group := r.Group("")
	group.Use(middlewares.JWTAuthMiddleware.MiddlewareFunc())
	group.GET("", controller.List)
	group.GET("/:id", controller.Retrieve)
	group.POST("", controller.Create)
	group.PUT("/:id", controller.Update)
	group.DELETE("/:id", controller.Destroy)

}
