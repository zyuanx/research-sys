package middleware

import (
	"gin-research-sys/internal/conf"
	"github.com/gin-gonic/gin"
	"log"
)

func Privilege() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取用户的角色
		sub := "admin"
		//获取请求的URI
		obj := c.Request.URL.RequestURI()
		obj = c.Request.URL.Path
		//获取请求方法
		act := c.Request.Method

		log.Println(obj, act, sub)
		//判断策略中是否存在
		result, err := conf.Enforcer.Enforce(sub, obj, act)
		if err != nil {
			c.JSON(400, gin.H{"code": 400, "message": err.Error()})
			c.Abort()
			return
		}
		if result {
			c.Next()
		} else {
			c.JSON(200, gin.H{"code": 200, "message": "access denied"})
			c.Abort()
		}
	}
}
