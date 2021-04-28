package middlewares

import (
	"gin-research-sys/pkg/global"
	"github.com/gin-gonic/gin"
	"log"
)

func Privilege() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userName = c.GetHeader("userName")
		if userName == "" {
			c.JSON(200, gin.H{"code": 200, "message": "header miss userName"})
			c.Abort()
			return
		}
		path := c.Request.URL.Path
		method := c.Request.Method
		cacheName := userName + path + method
		// 从缓存中读取&判断
		entry, err := global.Cache.Get(cacheName)
		if err == nil && entry != nil {
			if string(entry) == "true" {
				c.Next()
			} else {
				c.JSON(200, gin.H{"code": 200, "message": "access denied"})
				c.Abort()
				return
			}
		} else {
			// 从数据库中读取&判断
			//记录日志
			global.Enforcer.EnableLog(true)
			// 加载策略规则
			err := global.Enforcer.LoadPolicy()
			if err != nil {
				c.JSON(200, gin.H{"code": 200, "message": "loadPolicy error"})
				log.Println("loadPolicy error")
				panic(err)
			}
			// 验证策略规则
			log.Println(userName, path, method)
			result, err := global.Enforcer.Enforce(userName, path, method)
			if err != nil {
				c.JSON(200, gin.H{"code": 200, "message": "No permission found"})
				c.Abort()
				return
			}
			if !result {
				// 添加到缓存中
				err := global.Cache.Set(cacheName, []byte("false"))
				if err != nil {
					c.JSON(200, gin.H{"code": 200, "message": "add cache error"})
					c.Abort()
					return
				}
				c.JSON(200, gin.H{"code": 200, "message": "access denied"})
				c.Abort()
				return
			} else {
				err := global.Cache.Set(cacheName, []byte("true"))
				if err != nil {
					c.JSON(200, gin.H{"code": 200, "message": "add cache error"})
					c.Abort()
					return
				}
			}
			c.Next()
		}
	}
}
