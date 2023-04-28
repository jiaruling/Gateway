package tcp_proxy_middleware

import (
	"fmt"
	"strings"

	"github.com/jiaruling/Gateway/dao"
	"github.com/jiaruling/Gateway/global"
	ts "github.com/jiaruling/Gateway/proxy/tcp/server"
	"github.com/jiaruling/Gateway/public"
)

func TCPFlowLimitMiddleware() func(c *ts.TcpSliceRouterContext) {
	return func(c *ts.TcpSliceRouterContext) {
		serverInterface := c.Get("service")
		if serverInterface == nil {
			c.Conn.Write([]byte("get service empty"))
			c.Abort()
			return
		}
		serviceDetail := serverInterface.(*dao.ServiceDetail)

		if serviceDetail.AccessControl.ServiceFlowLimit != 0 {
			serviceLimiter, err := public.FlowLimiterHandler.GetLimiter(
				global.FlowServicePrefix+serviceDetail.Info.ServiceName,
				float64(serviceDetail.AccessControl.ServiceFlowLimit))
			if err != nil {
				c.Conn.Write([]byte(err.Error()))
				c.Abort()
				return
			}
			if !serviceLimiter.Allow() {
				c.Conn.Write([]byte(fmt.Sprintf("service flow limit %v", serviceDetail.AccessControl.ServiceFlowLimit)))
				c.Abort()
				return
			}
		}

		splits := strings.Split(c.Conn.RemoteAddr().String(), ":")
		clientIP := ""
		if len(splits) == 2 {
			clientIP = splits[0]
		}
		if serviceDetail.AccessControl.ClientIPFlowLimit > 0 {
			clientLimiter, err := public.FlowLimiterHandler.GetLimiter(
				global.FlowServicePrefix+serviceDetail.Info.ServiceName+"_"+clientIP,
				float64(serviceDetail.AccessControl.ClientIPFlowLimit))
			if err != nil {
				c.Conn.Write([]byte(err.Error()))
				c.Abort()
				return
			}
			if !clientLimiter.Allow() {
				c.Conn.Write([]byte(fmt.Sprintf("%v flow limit %v", clientIP, serviceDetail.AccessControl.ClientIPFlowLimit)))
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
