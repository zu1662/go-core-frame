package middleware

import (
	"bytes"
	"encoding/json"
	"go-core-frame/global"
	"go-core-frame/models"
	"go-core-frame/pkg/app"
	"go-core-frame/pkg/config"
	"go-core-frame/utils"
	"io/ioutil"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
)

// SetLogger 日志记录中间件
func SetLogger() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 获取 Request Body 信息
		var bodyParams []byte
		bodyParams, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			global.Logger.Fatal(" Read body from request error: ", err)
		} else if bodyParams != nil {
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyParams))
		}

		// 把 后续处理的 response body 信息储存
		writer := responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = writer

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
		reqURI := utils.GetAPIPath(c.FullPath(), "")

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		// 日志格式
		// fmt.Printf("%s [INFO] %s %s %3d %13v %15s \r\n",
		// 	startTime.Format("2006-01-02 15:04:05"),
		// 	reqMethod,
		// 	reqURI,
		// 	statusCode,
		// 	latencyTime,
		// 	clientIP,
		// )
		// 打印日志到文件
		global.Logger.Info(statusCode, latencyTime, clientIP, reqMethod, reqURI)

		// 操作日志/登录日志 保存到数据库
		if c.Request.Method != "GET" && c.Request.Method != "OPTIONS" {
			// 如果配置开启了保存在数据库
			if config.LoggerConfig.EnabledDB {
				LoggerToDB(c, clientIP, statusCode, reqURI, reqMethod, latencyTime, writer, string(bodyParams))
			}
		}
	}
}

// LoggerToDB 保存请求日志到数据库
func LoggerToDB(c *gin.Context, clientIP string, statusCode int, reqURI string, reqMethod string, latencyTime time.Duration, writer responseBodyWriter, params string) {

	ipLocation := utils.GetLocation(clientIP)
	ua := user_agent.New(c.Request.Header.Get("User-Agent"))

	// 返回的 信息
	var response app.Response
	json.Unmarshal([]byte(writer.body.String()), &response)

	// 获取当前用户信息
	userClaims := utils.GetUserClaims(c)

	if strings.Contains(reqURI, "/login") {
		// 用户登录时，暂无 token 信息，设置的获取 username 信息
		username, ok := c.Get("username")
		if !ok {
			username = "_"
		}

		loginLog := models.LoginLog{}
		loginLog.UserName = username.(string)
		loginLog.IPAddress = clientIP
		loginLog.IPLocation = ipLocation
		loginLog.Browser, _ = ua.Browser()
		loginLog.OS = ua.OS()
		loginLog.Result = response.Msg
		loginLog.LoginTime = utils.GetCurrentTime()

		loginLog.Create()
	} else {
		operLog := models.OperLog{}
		operLog.IPAddress = clientIP
		operLog.IPLocation = ipLocation
		operLog.OperName = userClaims.Username
		operLog.Method = reqMethod
		operLog.Path = reqURI
		operLog.LatencyTime = (latencyTime).String()
		operLog.Browser, _ = ua.Browser()
		operLog.OS = ua.OS()
		operLog.Result = response.Msg
		operLog.Response = writer.body.String()
		operLog.OperTime = utils.GetCurrentTime()

		operLog.Params = params

		// 获取 API Title
		api := models.SysAPI{}
		api.Path = strings.TrimLeft(reqURI, "/v1")
		api.Method = reqMethod
		nowAPI, apiErr := api.GetAPI()

		if apiErr == nil {
			operLog.OperTitle = nowAPI.Name
		}

		operLog.Create()
	}

}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
