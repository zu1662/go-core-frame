package router

import (
	"go-core-frame/global"
	"go-core-frame/middleware"
	"go-core-frame/pkg/config"

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

	// 设置API权限
	Router.Use(middleware.APIAuth())

	// 自定义错误处理
	Router.Use(middleware.CustomError)

	// 自定义 Header 头信息
	Router.Use(middleware.Options)

	// 允许跨域 CORS 请求
	Router.Use(middleware.CORS)

	// 设置请求安全模式
	Router.Use(middleware.Secure)

	// 设置请求不缓存
	Router.Use(middleware.NoCache)

	// 方便统一添加路由组前缀 多服务器上线使用
	APIGroup := Router.Group(config.ApplicationConfig.APIVersion)
	InitSystemRouter(APIGroup)

	return Router
}
