package main

import (
	"os"
	"os/signal"
	"syscall"

	_ "github.com/jiaruling/Gateway/initial"
	"github.com/jiaruling/Gateway/router"
)

func main() {
	// 启动服务
	router.HttpServerRun()
	// 优雅退出
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	// 销毁服务
	router.HttpServerStop()
}
