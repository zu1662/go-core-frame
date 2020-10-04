package middleware

import (
	"fmt"
	"go-core-frame/global"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// CustomError 自定义 pinic 处理
// 此函数会在触发 pinic 之后， recover 到错误信息
// 并且在检测到 gin 的当前请求被取消时
// 通过处理 pinic 的信息，把当前的错误返回
func CustomError(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			if c.IsAborted() {
				c.Status(200)
			}

			switch errStr := err.(type) {
			case string:
				p := strings.Split(errStr, "#")
				panicName := p[0]
				code := p[1]
				msg := p[2]

				if len(p) == 3 && panicName == "CustomError" {
					statusCode, e := strconv.Atoi(code)
					if e != nil {
						break
					}
					c.Status(statusCode)
					fmt.Println(
						time.Now().Format("2006-01-02 15:04:05"),
						"[ERROR]",
						c.Request.Method,
						c.Request.URL,
						statusCode,
						c.Request.RequestURI,
						c.ClientIP(),
						msg,
					)
					global.Logger.Error(" panic error :", err)
					c.JSON(http.StatusOK, gin.H{
						"code": statusCode,
						"msg":  msg,
					})
					break
				}
			default:
				panic(err)
			}
		}
	}()
	c.Next()
}
