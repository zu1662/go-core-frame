package router

import (
	"go-core-frame/global"

	"github.com/gin-gonic/gin"
)

// InitRouter Router initiazition
func InitRouter() *gin.Engine {
	var r *gin.Engine

	// 返回全局唯一gin实例
	if global.GinEngine == nil {
		r = gin.New()
	} else {
		r = global.GinEngine
	}

	// 注册系统内路由
	initSysRouter(r)

	return r
}
