package tcp_proxy_middleware

import (
	"fmt"
	"strings"

	"github.com/jiaruling/Gateway/dao"
	ts "github.com/jiaruling/Gateway/proxy/tcp/server"
	"github.com/jiaruling/Gateway/public"
)

// 匹配接入方式 基于请求信息
func TCPWhiteListMiddleware() func(c *ts.TcpSliceRouterContext) {
	return func(c *ts.TcpSliceRouterContext) {
		serverInterface := c.Get("service")
		if serverInterface == nil {
			c.Conn.Write([]byte("get service empty"))
			c.Abort()
			return
		}
		serviceDetail := serverInterface.(*dao.ServiceDetail)
		splits := strings.Split(c.Conn.RemoteAddr().String(), ":")
		clientIP := ""
		if len(splits) == 2 {
			clientIP = splits[0]
		}

		iplist := []string{}
		if serviceDetail.AccessControl.WhiteList != "" {
			iplist = strings.Split(serviceDetail.AccessControl.WhiteList, ",")
		}
		if serviceDetail.AccessControl.OpenAuth == 1 && len(iplist) > 0 {
			if !public.InStringSlice(iplist, clientIP) {
				c.Conn.Write([]byte(fmt.Sprintf("%s not in white ip list", clientIP)))
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
