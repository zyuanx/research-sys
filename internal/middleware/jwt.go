package middleware

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zyuanx/research-sys/internal/pkg/config"
	"github.com/zyuanx/research-sys/internal/pkg/constant"
	"github.com/zyuanx/research-sys/internal/pkg/errors"
	"github.com/zyuanx/research-sys/internal/pkg/errors/ecode"
	"github.com/zyuanx/research-sys/internal/pkg/jwt"
	"github.com/zyuanx/research-sys/internal/pkg/response"
)

const authorizationHeader = "Authorization"

// AuthToken 鉴权，验证用户token是否有效
func AuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := getJwtFromHeader(c)
		if err != nil {
			response.JSON(c, errors.Wrap(err, ecode.RequireAuthErr, "invalid token"), nil)
			c.Abort()
			return
		}
		// 验证token是否正确
		claims, err := jwt.ParseToken(token, config.GlobalConfig.JwtSecret)
		if err != nil {
			response.JSON(c, errors.Wrap(err, ecode.RequireAuthErr, "invalid token"), nil)
			c.Abort()
			return
		}
		c.Set(constant.UserID, claims.UserId)
		c.Next()
	}
}

func getJwtFromHeader(c *gin.Context) (string, error) {
	aHeader := c.Request.Header.Get(authorizationHeader)
	if len(aHeader) == 0 {
		return "", fmt.Errorf("token is empty")
	}
	strList := strings.SplitN(aHeader, " ", 2)
	if len(strList) != 2 || strList[0] != "Bearer" {
		return "", fmt.Errorf("token 不符合规则")
	}
	return strList[1], nil
}
