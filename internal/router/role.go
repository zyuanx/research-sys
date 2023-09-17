package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zyuanx/research-sys/internal/middleware"
)

func RegisterRoleRouter(r *gin.RouterGroup) {

	group := r.Group("")
	group.Use(middleware.AuthToken())
	group.GET("", c.RoleList)
	group.GET("/:id", c.RoleRetrieve)
	group.POST("", c.RoleCreate)
	group.PUT("/:id", c.RoleUpdate)
	group.DELETE("/:id", c.RoleDelete)

}
