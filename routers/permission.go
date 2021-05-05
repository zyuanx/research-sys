package routers

import (
	"gin-research-sys/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterPermissionRouter(r *gin.RouterGroup) {

	controller := controllers.NewPermissionController()

	role := r.Group("")
	//role.Use(middlewares.JWTAuthMiddleware.MiddlewareFunc())
	role.GET("", controller.List)
	role.GET("/:id", controller.Retrieve)
	role.POST("", controller.Create)
	role.PUT("/:id", controller.Update)

}
