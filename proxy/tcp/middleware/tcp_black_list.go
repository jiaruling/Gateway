package tcp_proxy_middleware

import (
	"fmt"
	"strings"

	"github.com/jiaruling/Gateway/dao"
	ts "github.com/jiaruling/Gateway/proxy/tcp/server"
	"github.com/jiaruling/Gateway/public"
)

// 匹配接入方式 基于请求信息
func TCPBlackListMiddleware() func(c *ts.TcpSliceRouterContext) {
	return func(c *ts.TcpSliceRouterContext) {
		serverInterface := c.Get("service")
		if serverInterface == nil {
			c.Conn.Write([]byte("get service empty"))
			c.Abort()
			return
		}
		serviceDetail := serverInterface.(*dao.ServiceDetail)

		whileIpList := []string{}
		if serviceDetail.AccessControl.WhiteList != "" {
			whileIpList = strings.Split(serviceDetail.AccessControl.WhiteList, ",")
		}

		blackIpList := []string{}
		if serviceDetail.AccessControl.BlackList != "" {
			blackIpList = strings.Split(serviceDetail.AccessControl.BlackList, ",")
		}

		splits := strings.Split(c.Conn.RemoteAddr().String(), ":")
		clientIP := ""
		if len(splits) == 2 {
			clientIP = splits[0]
		}
		if serviceDetail.AccessControl.OpenAuth == 1 && len(whileIpList) == 0 && len(blackIpList) > 0 {
			if public.InStringSlice(blackIpList, clientIP) {
				c.Conn.Write([]byte(fmt.Sprintf("%s in black ip list", clientIP)))
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
