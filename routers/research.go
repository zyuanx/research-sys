package routers

import (
	"gin-research-sys/controllers"
	"gin-research-sys/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterResearchRouter(r *gin.RouterGroup) {

	controller := controllers.NewResearchController()

	role := r.Group("")
	role.Use(middlewares.JWTAuthMiddleware.MiddlewareFunc())
	role.GET("/:id", controller.Retrieve)

}
