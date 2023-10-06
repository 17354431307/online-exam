package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	ERROR   = 7
	SUCCESS = 0
)

type Response struct {
	Code int
	Data any
	Msg  string
}

// Result 返回 response 的通用格式，
// 统一返回 200 的响应码代表着服务端已处理了，
// 具体真实的响应码是其内部 code 字段
func Result(code int, data any, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]any{}, message, c)
}

func FailWithDetailed(data any, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]any{}, message, c)
}
