package middleware

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jiaruling/Gateway/dao"
	m "github.com/jiaruling/Gateway/middleware"
	"github.com/jiaruling/Gateway/public"
)

// 匹配接入方式 基于请求信息
func HTTPBlackListMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		serverInterface, ok := c.Get("service")
		if !ok {
			m.ResponseError(c, 2001, errors.New("service not found"))
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
		if serviceDetail.AccessControl.OpenAuth == 1 && len(whileIpList) == 0 && len(blackIpList) > 0 {
			if public.InStringSlice(blackIpList, c.ClientIP()) {
				m.ResponseError(c, 3001, errors.New(fmt.Sprintf("%s in black ip list", c.ClientIP())))
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
