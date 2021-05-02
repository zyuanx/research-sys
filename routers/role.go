package routers

import (
	"gin-research-sys/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoleRouter(r *gin.RouterGroup) {

	controller := controllers.NewRoleController()

	role := r.Group("")
	//role.Use(middlewares.JWTAuthMiddleware.MiddlewareFunc())
	role.GET("", controller.List)
	role.POST("", controller.Create)
	role.GET("/:id", controller.Retrieve)

}
