package routers

import (
	"gin-research-sys/controllers"
	"gin-research-sys/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoleRouter(r *gin.RouterGroup) {

	controller := controllers.NewRoleController()

	role := r.Group("")
	role.Use(middlewares.JWTAuthMiddleware.MiddlewareFunc())
	role.GET("", controller.List)
	role.GET("/:id", controller.Retrieve)
	role.POST("", controller.Create)
	role.PUT("/:id", controller.Update)
	role.DELETE("/:id", controller.Destroy)

}
