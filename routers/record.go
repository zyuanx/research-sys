package routers

import (
	"gin-research-sys/controllers"
	"gin-research-sys/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRecordRouter(r *gin.RouterGroup) {

	controller := controllers.NewRecordController()

	group := r.Group("")
	group.Use(middlewares.JWTAuthMiddleware.MiddlewareFunc())
	group.GET("", controller.List)
	group.GET("/:id", controller.Retrieve)
	group.POST("", controller.Create)
}
