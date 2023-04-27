package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/jiaruling/Gateway/dao"
	_ "github.com/jiaruling/Gateway/initial"
	http_router "github.com/jiaruling/Gateway/proxy/http/router"
	tcp_router "github.com/jiaruling/Gateway/proxy/tcp/router"
	"github.com/jiaruling/Gateway/router"
)

//config ./conf/prod/ 对应配置文件夹

var (
	config = flag.String("config", "", "input config file like ./conf/dev/")
)

func main() {
	// 解析命令行参数
	flag.Parse()
	if *config == "" {
		flag.Usage()
		os.Exit(1)
	}
	dao.ServiceManagerHandler.LoadOnce() // 一次性加载所有服务至内存
	dao.AppManagerHandler.LoadOnce()     // 一次性加载所有租户至内存
	go router.HttpServerRun()            // 启动管理后台
	go http_router.HttpServerRun()       // 启动http反向代理
	go http_router.HttpsServerRun()      // 启动https反向代理
	go tcp_router.TcpServerRun()         // 启动tcp反向代理
	// 优雅退出
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	router.HttpServerStop()
	http_router.HttpServerStop()
	http_router.HttpsServerStop()
	tcp_router.TcpServerStop()
}
