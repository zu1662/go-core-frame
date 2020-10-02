package middleware

import (
	"fmt"
	"go-core-frame/global"
	"go-core-frame/models"
	"go-core-frame/pkg/config"
	"go-core-frame/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
)

// SetLogger 日志记录中间件
func SetLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqURI := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		// 日志格式
		fmt.Printf("%s [INFO] %s %s %3d %13v %15s \r\n",
			startTime.Format("2006-01-02 15:04:05"),
			reqMethod,
			reqURI,
			statusCode,
			latencyTime,
			clientIP,
		)

		// 打印日志到文件
		global.Logger.Info(statusCode, latencyTime, clientIP, reqMethod, reqURI)

		// 操作日志/登录日志 保存到数据库
		if c.Request.Method != "GET" && c.Request.Method != "OPTIONS" {
			// 如果配置开启了保存在数据库
			if config.LoggerConfig.EnabledDB {
				LoggerToDB(c, clientIP, statusCode, reqURI, reqMethod, latencyTime)
			}
		}
	}
}

// LoggerToDB 保存请求日志到数据库
func LoggerToDB(c *gin.Context, clientIP string, statusCode int, reqURI string, reqMethod string, latencyTime time.Duration) {
	api := models.API{}
	api.Path = reqURI
	api.Method = reqMethod
	apiList, _ := api.Get()

	ipLocation := utils.GetLocation(clientIP)
	ua := user_agent.New(c.Request.Header.Get("User-Agent"))

	if reqURI == "/login" {
		loginLog := models.LoginLog{}
		loginLog.UserName = "Test"
		loginLog.IPAddress = clientIP
		loginLog.IPLocation = ipLocation
		loginLog.Browser, _ = ua.Browser()
		loginLog.OS = ua.OS()
		loginLog.Result = "OK"
		loginLog.LoginTime = utils.GetCurrentTime()

		loginLog.Create()
	} else {
		operLog := models.OperLog{}
		operLog.IPAddress = clientIP
		operLog.IPLocation = ipLocation
		operLog.OperName = "Test"
		operLog.Method = reqMethod
		operLog.Path = reqURI
		operLog.LatencyTime = (latencyTime).String()
		operLog.Browser, _ = ua.Browser()
		operLog.OS = ua.OS()
		operLog.Result = "OK"
		operLog.OperTime = utils.GetCurrentTime()

		body, _ := c.Get("body")
		operLog.Params, _ = utils.StructToJsonStr(body)

		if len(apiList) > 0 {
			operLog.OperTitle = apiList[0].Name
		}

		operLog.Create()
	}

}
