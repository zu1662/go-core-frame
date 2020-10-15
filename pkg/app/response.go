package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// OK 成功返回
func OK(c *gin.Context, msg string, data interface{}) {
	var res Response
	res.Code = 1
	res.Data = data
	if msg != "" {
		res.Msg = msg
	}
	c.JSON(http.StatusOK, res)
}

// Error 失败返回
func Error(c *gin.Context, code int, err error, msg string) {
	var res Response
	res.Code = code
	// 默认为错误信息
	res.Msg = err.Error()
	if msg != "" {
		res.Msg = msg
	}
	c.JSON(http.StatusOK, res)
}

// WithCode 自定义 responese Code 返回
func WithCode(c *gin.Context, code int, err error, msg string) {
	var res Response
	res.Code = code
	// 默认为错误信息
	res.Msg = err.Error()
	if msg != "" {
		res.Msg = msg
	}

	c.JSON(code, res)
}

// Custom 自定义返回 兼容函数
func Custom(c *gin.Context, data gin.H) {
	c.JSON(http.StatusOK, data)
}
