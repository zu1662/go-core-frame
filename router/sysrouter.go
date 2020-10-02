package router

import (
	"go-core-frame/api/system"

	"github.com/gin-gonic/gin"
)

func initSysRouter(r *gin.Engine) {
	g := r.Group("")

	// 一些基础路由，可用于测试
	sysBaseRouter(g)
}

func sysBaseRouter(r *gin.RouterGroup) {
	r.GET("/", system.HelloWorld)

	r.POST("/sysuser/:userId", system.GetSysUser)
}
