package router

import (
	"context"
	"fmt"
	"log"

	"github.com/jiaruling/Gateway/dao"
	tm "github.com/jiaruling/Gateway/proxy/tcp/middleware"
	ts "github.com/jiaruling/Gateway/proxy/tcp/server"
	"github.com/jiaruling/Gateway/reverse_proxy"
)

var tcpServerList = []*ts.TcpServer{}

type tcpHandler struct {
}

func TcpServerRun() {
	serviceList := dao.ServiceManagerHandler.GetTcpServiceList()
	for _, serviceItem := range serviceList {
		tempItem := serviceItem
		go func(serviceDetail *dao.ServiceDetail) {
			addr := fmt.Sprintf(":%d", serviceDetail.TCPRule.Port)
			rb, err := dao.LoadBalancerHandler.GetLoadBalancer(serviceDetail)
			if err != nil {
				log.Fatalf(" [INFO] GetTcpLoadBalancer %v err:%v\n", addr, err)
				return
			}

			//构建路由及设置中间件
			router := ts.NewTcpSliceRouter()
			router.Group("/").Use(
				tm.TCPFlowCountMiddleware(),
				tm.TCPFlowLimitMiddleware(),
				tm.TCPWhiteListMiddleware(),
				tm.TCPBlackListMiddleware(),
			)

			//构建回调handler
			routerHandler := ts.NewTcpSliceRouterHandler(
				func(c *ts.TcpSliceRouterContext) ts.TCPHandler {
					return reverse_proxy.NewTcpLoadBalanceReverseProxy(c, rb)
				}, router)

			baseCtx := context.WithValue(context.Background(), "service", serviceDetail)
			tcpServer := &ts.TcpServer{
				Addr:    addr,
				Handler: routerHandler,
				BaseCtx: baseCtx,
			}
			tcpServerList = append(tcpServerList, tcpServer)
			log.Printf(" [INFO] tcp_proxy_run %v\n", addr)
			if err := tcpServer.ListenAndServe(); err != nil && err != ts.ErrServerClosed {
				log.Fatalf(" [INFO] tcp_proxy_run %v err:%v\n", addr, err)
			}
		}(tempItem)
	}
}

func TcpServerStop() {
	for _, tcpServer := range tcpServerList {
		tcpServer.Close()
		log.Printf(" [INFO] tcp_proxy_stop %v stopped\n", tcpServer.Addr)
	}
}
