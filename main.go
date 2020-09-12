package main

import (
	"context"
	"fmt"
	"go-core-frame/pkg/config"
	"go-core-frame/pkg/logger"
	"go-core-frame/router"
	"go-core-frame/utils"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	run()
}

func run() {
	// 配置初始化
	setup()
	if config.ApplicationConfig.Mode == string(utils.ModeProd) {
		gin.SetMode(gin.ReleaseMode)
	}

	r := router.InitRouter()
	server := &http.Server{
		Addr:    config.ApplicationConfig.Host + ":" + config.ApplicationConfig.Port,
		Handler: r,
	}

	go func() {
		server.ListenAndServe()
	}()

	content, _ := ioutil.ReadFile("./static/copyright.txt")
	fmt.Println(utils.Red(string(content)))

	fmt.Println(utils.Green("Server run at:"))
	fmt.Printf("-  Local:   http://localhost:%s/ \r\n", config.ApplicationConfig.Port)
	fmt.Printf("-  Network: http://%s:%s/ \r\n", utils.GetLocaHonst(), config.ApplicationConfig.Port)
	fmt.Println(utils.Green("Swagger run at:"))
	fmt.Printf("-  Local:   http://localhost:%s/swagger/index.html \r\n", config.ApplicationConfig.Port)
	fmt.Printf("-  Network: http://%s:%s/swagger/index.html \r\n", utils.GetLocaHonst(), config.ApplicationConfig.Port)
	fmt.Printf("%s Enter Control + C Shutdown Server \r\n", utils.GetCurrentTimeStr())

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Printf("%s Shutdown Server ... \r\n", utils.GetCurrentTimeStr())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("%s Server Exited ... \r\n", utils.GetCurrentTimeStr())
	}
}

func setup() {
	// 配置初始化
	config.Setup("./config/settings.yml")

	// 日志初始化
	logger.Setup()
}
