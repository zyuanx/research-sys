package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Res struct {
	code int
	data gin.H
	message  string
}

func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, message string) {
	ctx.JSON(httpStatus, gin.H{"code": code, "data": data, "message": message})
}

func Success(ctx *gin.Context, code int, data gin.H, message string) {
	Response(ctx, http.StatusOK, code, data, message)
}
func Fail(ctx *gin.Context, code int, data gin.H, message string) {
	Response(ctx, http.StatusOK, code, data, message)
}
