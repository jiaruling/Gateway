package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/jiaruling/Gateway/dao"
	_ "github.com/jiaruling/Gateway/initial"
	hr "github.com/jiaruling/Gateway/proxy/http/router"
	"github.com/jiaruling/Gateway/router"
)

//endpoint dashboard后台管理  server代理服务器
//config ./conf/prod/ 对应配置文件夹

var (
	endpoint = flag.String("endpoint", "", "input endpoint dashboard or server")
	config   = flag.String("config", "", "input config file like ./conf/dev/")
)

func main() {
	// 解析命令行参数
	flag.Parse()
	if *endpoint == "" {
		flag.Usage()
		os.Exit(1)
	}
	if *config == "" {
		flag.Usage()
		os.Exit(1)
	}
	if *endpoint == "dashboard" {
		// 启动服务
		router.HttpServerRun()
		// 优雅退出
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		// 销毁服务
		router.HttpServerStop()
	} else {
		dao.ServiceManagerHandler.LoadOnce()
		dao.AppManagerHandler.LoadOnce()

		go hr.HttpServerRun()
		go hr.HttpsServerRun()

		// 优雅退出
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
		<-quit
	}
}
