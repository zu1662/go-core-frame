package router

import (
	"go-core-frame/api/system"
	"go-core-frame/middleware"

	// Swagger doc path
	_ "go-core-frame/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// InitSystemRouter 初始化路由
func InitSystemRouter(Router *gin.RouterGroup) {
	// 基础路由
	initBaseRouter(Router)

	// 日志路由
	initLogRouter(Router)
	// 用户路由
	initUserRouter(Router)
	// 部门路由
	initDeptRouter(Router)
	// 岗位路由
	initPostRouter(Router)
	// 菜单路由
	initMenuRouter(Router)
	// 角色路由
	initRoleRouter(Router)
	// 角色菜单路由
	initRoleMenuRouter(Router)
}

// InitBaseRouter 基础路由，不需要鉴权
func initBaseRouter(Router *gin.RouterGroup) {
	Router.GET("/", system.HelloWorld)
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
		APIRouter.GET("/info/:userId", system.GetUserDetail)
		APIRouter.GET("/list", system.GetUserList)
		APIRouter.PUT("/update", system.UpdateUser)
		APIRouter.DELETE("/delete/:userId", system.DeleteUser)
		APIRouter.POST("/add", system.InsertUser)
	}
}

// initDeptRouter 部门路由
func initDeptRouter(Router *gin.RouterGroup) {
	APIRouter := Router.Group("dept").
		Use(middleware.JWTAuth())
	{
		APIRouter.GET("/info/:deptId", system.GetDeptDetail)
		APIRouter.GET("/tree", system.GetDeptTree)
		APIRouter.PUT("/update", system.UpdateDept)
		APIRouter.DELETE("/delete/:deptId", system.DeleteDept)
		APIRouter.POST("/add", system.InsertDept)
	}
}

// initPostRouter 岗位路由
func initPostRouter(Router *gin.RouterGroup) {
	APIRouter := Router.Group("post").
		Use(middleware.JWTAuth())
	{
		APIRouter.GET("/info/:postId", system.GetPostDetail)
		APIRouter.GET("/list", system.GetPostList)
		APIRouter.PUT("/update", system.UpdatePost)
		APIRouter.DELETE("/delete/:postId", system.DeletePost)
		APIRouter.POST("/add", system.InsertPost)
	}
}

// initPostRouter 岗位路由
func initMenuRouter(Router *gin.RouterGroup) {
	APIRouter := Router.Group("menu").
		Use(middleware.JWTAuth())
	{
		APIRouter.GET("/info/:menuId", system.GetMenuDetail)
		APIRouter.GET("/tree", system.GetMenuTree)
		APIRouter.PUT("/update", system.UpdateMenu)
		APIRouter.DELETE("/delete/:menuId", system.DeleteMenu)
		APIRouter.POST("/add", system.InsertMenu)
	}
}

// initRoleRouter 角色路由
func initRoleRouter(Router *gin.RouterGroup) {
	APIRouter := Router.Group("role").
		Use(middleware.JWTAuth())
	{
		APIRouter.GET("/info/:roleId", system.GetRoleDetail)
		APIRouter.GET("/list", system.GetRoleList)
		APIRouter.PUT("/update", system.UpdateRole)
		APIRouter.DELETE("/delete/:roleId", system.DeleteRole)
		APIRouter.POST("/add", system.InsertRole)
	}
}

// initRoleMenuRouter 角色菜单路由
func initRoleMenuRouter(Router *gin.RouterGroup) {
	APIRouter := Router.Group("rolemenu").
		Use(middleware.JWTAuth())
	{
		APIRouter.GET("/list", system.GetRoleMenu)
		APIRouter.POST("/update", system.UpdateRoleMenu)
	}
}
