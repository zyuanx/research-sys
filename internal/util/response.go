package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Code    int      `json:"code" example:"200"`
	Message string   `json:"message" example:"响应信息"`
	Data    struct{} `json:"data" `
}

func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	ctx.JSON(httpStatus, gin.H{"code": code, "data": data, "message": msg})
}

func Success(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 200, data, msg)
}
func Fail(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 400, data, msg)
}

type BaseData struct {
	Id        uint   `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type PaginationRes struct {
	Size    uint        `json:"size"`
	Page    uint        `json:"page"`
	Results interface{} `json:"results"`
	Total   uint        `json:"total"`
}
