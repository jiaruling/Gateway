package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jiaruling/Gateway/dao"
	m "github.com/jiaruling/Gateway/middleware"
)

// done: 匹配服务
func HTTPAccessModeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		service, err := dao.ServiceManagerHandler.HTTPAccessMode(c)
		if err != nil {
			m.ResponseError(c, 1001, err)
			c.Abort()
			return
		}
		//fmt.Println("matched service",public.Obj2Json(service))
		c.Set("service", service)
		c.Next()
	}
}
