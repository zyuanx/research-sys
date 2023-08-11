package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zyuanx/research-sys/internal/pkg/constant"
	"github.com/zyuanx/research-sys/internal/pkg/errors"
	"github.com/zyuanx/research-sys/internal/pkg/errors/ecode"
)

type ApiResponse struct {
	RequestId string      `json:"request_id"`     // 请求的唯一ID
	ErrCode   int         `json:"err_code"`       // 错误码，0表示无错误
	Message   string      `json:"message"`        // 提示信息
	Data      interface{} `json:"data,omitempty"` // 响应数据，一般从这里前端从这个里面取出数据展示
}

// JSON 发送json格式的数据
func JSON(c *gin.Context, err error, data interface{}) {
	errCode, message := errors.DecodeErr(err)
	var httpStatus int
	if errCode != ecode.Success {
		httpStatus = http.StatusBadRequest
	} else {
		httpStatus = http.StatusOK
	}
	c.JSON(httpStatus, ApiResponse{
		RequestId: c.GetString(constant.RequestId),
		ErrCode:   errCode,
		Message:   message,
		Data:      data,
	})
}

// func Success(ctx *gin.Context, data gin.H, msg string) {
// 	JSON(ctx, http.StatusOK, 200, data, msg)
// }
// func Fail(ctx *gin.Context, data gin.H, msg string) {
// 	JSON(ctx, http.StatusOK, 400, data, msg)
// }
