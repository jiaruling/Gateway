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
		dao.ServiceManagerHandler.LoadOnce() // 一次性加载所有服务至内存
		dao.AppManagerHandler.LoadOnce()     // 一次性加载所有租户至内存

		go hr.HttpServerRun()  // 启动http反向代理
		go hr.HttpsServerRun() // 启动https反向代理

		// 优雅退出
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
		<-quit
	}
}
