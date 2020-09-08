package main

import (
	"context"
	"fmt"
	"go-core-frame/router"
	"go-core-frame/tools"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	run()
}

func run() {
	// if false {
	// 	gin.SetMode(gin.ReleaseMode)
	// }

	r := router.InitRouter()
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		server.ListenAndServe()
	}()

	content, _ := ioutil.ReadFile("./static/copyright.txt")
	fmt.Println(tools.Red(string(content)))

	fmt.Println(tools.Green("Server run at:"))
	fmt.Printf("-  Local:   http://localhost:%s/ \r\n", "8080")
	fmt.Printf("-  Network: http://%s:%s/ \r\n", tools.GetLocaHonst(), "8080")
	fmt.Println(tools.Green("Swagger run at:"))
	fmt.Printf("-  Local:   http://localhost:%s/swagger/index.html \r\n", "8080")
	fmt.Printf("-  Network: http://%s:%s/swagger/index.html \r\n", tools.GetLocaHonst(), "8080")
	fmt.Printf("%s Enter Control + C Shutdown Server \r\n", tools.GetCurrentTimeStr())

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Printf("%s Shutdown Server ... \r\n", tools.GetCurrentTimeStr())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("%s Server Exited ... \r\n", tools.GetCurrentTimeStr())
	}
}
