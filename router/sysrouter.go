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
	// 字典路由
	initDicRouter(Router)
	// API路由
	initAPIRouter(Router)
	// 角色接口路由
	initRoleAPIRouter(Router)
}

// InitBaseRouter 基础路由，不需要鉴权
func initBaseRouter(Router *gin.RouterGroup) {
	Router.GET("/", system.HelloWorld)
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	BaseRouter := Router.Group("base")
	{
		BaseRouter.POST("/login", system.Login)
		BaseRouter.GET("/captcha", system.GetCaptcha)
		BaseRouter.POST("/logout", system.Logout)
		BaseRouter.
			Use(middleware.JWTAuth()).
			Use(middleware.APIAuth()).
			GET("/userinfo", system.GetUserInfo)
	}
}

// InitLogRouter 操作日志路由
func initLogRouter(Router *gin.RouterGroup) {
	APIRouter := Router.Group("log").
		Use(middleware.JWTAuth()).
		Use(middleware.APIAuth())
	{
		APIRouter.GET("/loginloginfo", system.GetLoginLogInfo)
		APIRouter.GET("/loginloglist", system.GetLoginLogList)
		APIRouter.DELETE("/deleteloginlog/:logIds", system.DeleteLoginlog)
		APIRouter.DELETE("/cleanloginlog", system.CleanLoginlog)
		APIRouter.GET("/operloginfo", system.GetOperLogInfo)
		APIRouter.GET("/operloglist", system.GetOperLogList)
		APIRouter.DELETE("/deleteoperlog/:logIds", system.DeleteOperlog)
		APIRouter.DELETE("/cleanoperlog", system.CleanOperlog)
	}
}

// InitUserRouter 用户路由
func initUserRouter(Router *gin.RouterGroup) {
	APIRouter := Router.Group("user").
		Use(middleware.JWTAuth()).
		Use(middleware.APIAuth())
	{
		APIRouter.GET("/info/:userId", system.GetUserDetail)
		APIRouter.GET("/list", system.GetUserList)
		APIRouter.GET("/listall", system.GetUserAll)
		APIRouter.PUT("/update", system.UpdateUser)
		APIRouter.PUT("/resetpsw", system.ResetUserPsw)
		APIRouter.DELETE("/delete/:userId", system.DeleteUser)
		APIRouter.POST("/add", system.InsertUser)
	}
}

// initDeptRouter 部门路由
func initDeptRouter(Router *gin.RouterGroup) {
	APIRouter := Router.Group("dept").
		Use(middleware.JWTAuth()).
		Use(middleware.APIAuth())
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
		Use(middleware.JWTAuth()).
		Use(middleware.APIAuth())
	{
		APIRouter.GET("/info/:postId", system.GetPostDetail)
		APIRouter.GET("/list", system.GetPostList)
		APIRouter.GET("/listall", system.GetPostAll)
		APIRouter.PUT("/update", system.UpdatePost)
		APIRouter.DELETE("/delete/:postId", system.DeletePost)
		APIRouter.POST("/add", system.InsertPost)
	}
}

// initMenuRouter 菜单路由
func initMenuRouter(Router *gin.RouterGroup) {
	APIRouter := Router.Group("menu").
		Use(middleware.JWTAuth()).
		Use(middleware.APIAuth())
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
		Use(middleware.JWTAuth()).
		Use(middleware.APIAuth())
	{
		APIRouter.GET("/info/:roleId", system.GetRoleDetail)
		APIRouter.GET("/list", system.GetRoleList)
		APIRouter.GET("/listall", system.GetRoleAll)
		APIRouter.PUT("/update", system.UpdateRole)
		APIRouter.DELETE("/delete/:roleId", system.DeleteRole)
		APIRouter.POST("/add", system.InsertRole)
	}
}

// initRoleMenuRouter 角色菜单路由
func initRoleMenuRouter(Router *gin.RouterGroup) {
	APIRouter := Router.Group("rolemenu").
		Use(middleware.JWTAuth()).
		Use(middleware.APIAuth())
	{
		APIRouter.GET("/list", system.GetRoleMenu)
		APIRouter.POST("/update", system.UpdateRoleMenu)
	}
}

// initDicRouter 部门路由
func initDicRouter(Router *gin.RouterGroup) {
	APIRouter := Router.Group("dict").
		Use(middleware.JWTAuth()).
		Use(middleware.APIAuth())
	{
		APIRouter.GET("/dicttypelist", system.GetDictTypeList)
		APIRouter.GET("/dicttype/:dictTypeId", system.GetDictTypeDetail)
		APIRouter.GET("/dictmap", system.GetDictMap)
		APIRouter.PUT("/dicttypeupdate", system.UpdateDictType)
		APIRouter.DELETE("/dicttypedelete/:dictTypeId", system.DeleteDictType)
		APIRouter.POST("/dicttypeadd", system.InsertDictType)

		APIRouter.GET("/dictdatalist", system.GetDictDataList)
		APIRouter.GET("/dictdata/:dictDataId", system.GetDictDataDetail)
		APIRouter.PUT("/dictdataupdate", system.UpdateDictData)
		APIRouter.DELETE("/dictdatadelete/:dictDataId", system.DeleteDictData)
		APIRouter.POST("/dictdataadd", system.InsertDictData)
	}
}

// initAPIRouter 岗位路由
func initAPIRouter(Router *gin.RouterGroup) {
	APIRouter := Router.Group("interface").
		Use(middleware.JWTAuth()).
		Use(middleware.APIAuth())
	{
		APIRouter.GET("/info/:apiId", system.GetAPIDetail)
		APIRouter.GET("/tree", system.GetAPITree)
		APIRouter.PUT("/update", system.UpdateAPI)
		APIRouter.DELETE("/delete/:apiId", system.DeleteAPI)
		APIRouter.POST("/add", system.InsertAPI)
	}
}

// initRoleAPIRouter 角色接口路由
func initRoleAPIRouter(Router *gin.RouterGroup) {
	APIRouter := Router.Group("roleapi").
		Use(middleware.JWTAuth()).
		Use(middleware.APIAuth())
	{
		APIRouter.GET("/list", system.GetRoleAPI)
		APIRouter.POST("/update", system.UpdateRoleAPI)
	}
}
