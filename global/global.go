package global

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/gogf/gf/os/glog"
	"gorm.io/gorm"
)

// GinEngine 全局的gin
var GinEngine *gin.Engine

// Logger 全局业务日志
var Logger *glog.Logger

// RequestLogger 全局请求日志
var RequestLogger *glog.Logger

// DB 全局 Mysql gorm实例
var DB *gorm.DB

// Redis 全局 Redis 实例
var Redis *redis.Client
