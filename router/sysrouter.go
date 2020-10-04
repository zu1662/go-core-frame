package router

import (
	"go-core-frame/api/system"
	"go-core-frame/middleware"

	"github.com/gin-gonic/gin"
)

// InitSystemRouter 初始化路由
func InitSystemRouter(Router *gin.RouterGroup) {
	// 基础路由
	initBaseRouter(Router)

	// 用户路由
	initUserRouter(Router)
}

// InitBaseRouter 基础路由，不需要鉴权
func initBaseRouter(Router *gin.RouterGroup) {
	Router.GET("/", system.HelloWorld)
	Router.POST("/login", system.Login)
}

// InitUserRouter 用户路由
func initUserRouter(Router *gin.RouterGroup) {
	APIRouter := Router.Group("user").
		Use(middleware.JWTAuth())
	{
		APIRouter.GET("/:userId", system.GetSysUser)
	}
}
