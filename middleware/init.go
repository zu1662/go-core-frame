package middleware

import (
	"github.com/gin-gonic/gin"
)

// InitMiddleware 初始化中间件
func InitMiddleware(r *gin.Engine) {
	// 自定义日志处理
	r.Use(SetLogger())

	// 自定义错误处理
	r.Use(CustomError)
}
