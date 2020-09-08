package router

import (
	"fmt"
	"go-core-frame/api/v1/system"

	"github.com/gin-gonic/gin"
)

func initSysRouter(r *gin.Engine) {
	g := r.Group("")

	// 一些基础路由，可用于测试
	sysBaseRouter(g)
	fmt.Println("now is showing")
}

func sysBaseRouter(r *gin.RouterGroup) {
	r.GET("/", system.HelloWorld)

	r.GET("/ping", system.Ping)
}
