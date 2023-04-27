package middleware

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jiaruling/Gateway/dao"
	"github.com/jiaruling/Gateway/global"
	m "github.com/jiaruling/Gateway/middleware"
	"github.com/jiaruling/Gateway/public"
)

// done: 流量限制
func HTTPFlowLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		serverInterface, ok := c.Get("service")
		if !ok {
			m.ResponseError(c, 2001, errors.New("service not found"))
			c.Abort()
			return
		}
		serviceDetail := serverInterface.(*dao.ServiceDetail)

		// 服务端限流
		if serviceDetail.AccessControl.ServiceFlowLimit != 0 {
			serviceLimiter, err := public.FlowLimiterHandler.GetLimiter(
				global.FlowServicePrefix+serviceDetail.Info.ServiceName,
				float64(serviceDetail.AccessControl.ServiceFlowLimit))
			if err != nil {
				m.ResponseError(c, 5001, err)
				c.Abort()
				return
			}
			if !serviceLimiter.Allow() {
				m.ResponseError(c, 5002, errors.New(fmt.Sprintf("service flow limit %v", serviceDetail.AccessControl.ServiceFlowLimit)))
				c.Abort()
				return
			}
		}

		// 客户端限流
		if serviceDetail.AccessControl.ClientIPFlowLimit > 0 {
			clientLimiter, err := public.FlowLimiterHandler.GetLimiter(
				global.FlowServicePrefix+serviceDetail.Info.ServiceName+"_"+c.ClientIP(),
				float64(serviceDetail.AccessControl.ClientIPFlowLimit))
			if err != nil {
				m.ResponseError(c, 5003, err)
				c.Abort()
				return
			}
			if !clientLimiter.Allow() {
				m.ResponseError(c, 5002, errors.New(fmt.Sprintf("%v flow limit %v", c.ClientIP(), serviceDetail.AccessControl.ClientIPFlowLimit)))
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
