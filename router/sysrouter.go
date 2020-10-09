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

	// 日志路由
	initLogRouter(Router)
	// 用户路由
	initUserRouter(Router)
}

// InitBaseRouter 基础路由，不需要鉴权
func initBaseRouter(Router *gin.RouterGroup) {
	Router.GET("/", system.HelloWorld)

	BaseRouter := Router.Group("base")
	{
		BaseRouter.POST("/login", system.Login)
		BaseRouter.Use(middleware.JWTAuth()).POST("/logout", system.Logout)
		BaseRouter.Use(middleware.JWTAuth()).GET("/userinfo", system.GetUserInfo)
	}
}

// InitLogRouter 操作日志路由
func initLogRouter(Router *gin.RouterGroup) {
	APIRouter := Router.Group("log").
		Use(middleware.JWTAuth())
	{
		APIRouter.GET("/loginloginfo", system.GetLoginLogInfo)
		APIRouter.GET("/loginloglist", system.GetLoginLogList)
		APIRouter.GET("/operloginfo", system.GetOperLogInfo)
		APIRouter.GET("/operloglist", system.GetOperLogList)
	}
}

// InitUserRouter 用户路由
func initUserRouter(Router *gin.RouterGroup) {
	APIRouter := Router.Group("user").
		Use(middleware.JWTAuth())
	{
		APIRouter.GET("/userinfo/:userId", system.GetUserDetail)
		APIRouter.GET("/userlist", system.GetUserList)
		APIRouter.PUT("/userupdate", system.UpdateUser)
		APIRouter.DELETE("/userdelete/:userId", system.DeleteUser)
		APIRouter.POST("/useradd", system.InsertUser)
	}
}
