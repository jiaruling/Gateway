package middleware

import (
	"errors"

	"github.com/jiaruling/Gateway/dao"
	"github.com/jiaruling/Gateway/reverse_proxy"

	"github.com/gin-gonic/gin"
	m "github.com/jiaruling/Gateway/middleware"
)

// done: 反向代理
func HTTPReverseProxyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		serverInterface, ok := c.Get("service")
		if !ok {
			m.ResponseError(c, 2001, errors.New("service not found"))
			c.Abort()
			return
		}
		serviceDetail := serverInterface.(*dao.ServiceDetail)

		lb, err := dao.LoadBalancerHandler.GetLoadBalancer(serviceDetail)
		if err != nil {
			m.ResponseError(c, 2002, err)
			c.Abort()
			return
		}
		trans, err := dao.TransportorHandler.GetTrans(serviceDetail)
		if err != nil {
			m.ResponseError(c, 2003, err)
			c.Abort()
			return
		}

		proxy := reverse_proxy.NewLoadBalanceReverseProxy(c, lb, trans)
		proxy.ServeHTTP(c.Writer, c.Request)
		c.Abort()
		return
	}
}
