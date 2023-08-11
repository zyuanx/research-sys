package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/zyuanx/research-sys/internal/pkg/constant"
	"github.com/zyuanx/research-sys/tools"
)

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := tools.GenUUID()
		c.Header("X-Request-Id", requestId)

		// 设置requestId到context中，便于后面调用链的透传
		c.Set(constant.RequestId, requestId)
		c.Next()
	}
}
