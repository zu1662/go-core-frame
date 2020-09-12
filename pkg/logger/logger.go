package logger

import (
	"go-core-frame/global"
	"go-core-frame/pkg/config"

	"github.com/gogf/gf/os/glog"
)

// Logger Business 业务 日志
var Logger *glog.Logger

// RequestLogger 请求日志
var RequestLogger *glog.Logger

// Setup 日志 初始化
func Setup() {
	Logger = glog.New()
	// 日志路径
	Logger.SetPath(config.LoggerConfig.Path + "/bus")
	// 是否打印在控制台
	Logger.SetStdoutPrint(config.LoggerConfig.EnabledBus && config.LoggerConfig.Stdout)
	// 设置日志文件名称格式
	Logger.SetFile("bus-{Ymd}.log")
	// 设置日志级别
	Logger.SetLevelStr(config.LoggerConfig.Level)

	RequestLogger = glog.New()
	// 日志路径
	RequestLogger.SetPath(config.LoggerConfig.Path + "/request")
	// 是否打印在控制台
	RequestLogger.SetStdoutPrint(config.LoggerConfig.EnabledBus && config.LoggerConfig.Stdout)
	// 设置日志文件名称格式
	RequestLogger.SetFile("request-{Ymd}.log")
	// 设置日志级别
	RequestLogger.SetLevelStr(config.LoggerConfig.Level)

	Logger.Info("Logger init success!")

	global.Logger = Logger.Line()
	global.RequestLogger = RequestLogger.Line()
}
