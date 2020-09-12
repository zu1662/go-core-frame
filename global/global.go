package global

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/os/glog"
)

// GinEngine 全局的gin
var GinEngine *gin.Engine

// Logger 全局业务日志
var Logger *glog.Logger

// RequestLogger 全局请求日志
var RequestLogger *glog.Logger
