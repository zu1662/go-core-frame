package router

import (
	"go-core-frame/global"
	"go-core-frame/middleware"

	"github.com/gin-gonic/gin"
)

// InitRouter Router initiazition
func InitRouter() *gin.Engine {
	var Router *gin.Engine

	// 返回全局唯一gin实例
	if global.GinEngine == nil {
		Router = gin.New()
	} else {
		Router = global.GinEngine
	}

	// 注册中间件
	// 自定义日志处理
	Router.Use(middleware.SetLogger())

	// 自定义错误处理
	Router.Use(middleware.CustomError)

	// 方便统一添加路由组前缀 多服务器上线使用
	APIGroup := Router.Group("v1")
	InitSystemRouter(APIGroup)

	return Router
}
