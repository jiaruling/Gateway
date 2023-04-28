package tcp_proxy_middleware

import (
	"github.com/jiaruling/Gateway/dao"
	"github.com/jiaruling/Gateway/global"
	ts "github.com/jiaruling/Gateway/proxy/tcp/server"
	"github.com/jiaruling/Gateway/public"
)

func TCPFlowCountMiddleware() func(c *ts.TcpSliceRouterContext) {
	return func(c *ts.TcpSliceRouterContext) {
		serverInterface := c.Get("service")
		if serverInterface == nil {
			c.Conn.Write([]byte("get service empty"))
			c.Abort()
			return
		}
		serviceDetail := serverInterface.(*dao.ServiceDetail)

		//统计项 1 全站 2 服务
		totalCounter, err := public.FlowCounterHandler.GetCounter(global.FlowTotal)
		if err != nil {
			c.Conn.Write([]byte(err.Error()))
			c.Abort()
			return
		}
		totalCounter.Increase()

		serviceCounter, err := public.FlowCounterHandler.GetCounter(global.FlowServicePrefix + serviceDetail.Info.ServiceName)
		if err != nil {
			c.Conn.Write([]byte(err.Error()))
			c.Abort()
			return
		}
		serviceCounter.Increase()
		c.Next()
	}
}
