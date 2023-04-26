package middleware

import (
	"errors"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jiaruling/Gateway/dao"
	m "github.com/jiaruling/Gateway/middleware"
)

// done: url重写 例如：将/aaa/foo/bar重定向为/bbb/foo/bar，将/aaa/baz重定向为/bbb/baz
func HTTPUrlRewriteMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		serverInterface, ok := c.Get("service")
		if !ok {
			m.ResponseError(c, 2001, errors.New("service not found"))
			c.Abort()
			return
		}
		serviceDetail := serverInterface.(*dao.ServiceDetail)
		for _, item := range strings.Split(serviceDetail.HTTPRule.UrlRewrite, ",") {
			items := strings.Split(item, " ")
			if len(items) != 2 {
				continue
			}
			regexp, err := regexp.Compile(items[0])
			if err != nil {
				continue
			}
			replacePath := regexp.ReplaceAll([]byte(c.Request.URL.Path), []byte(items[1]))
			c.Request.URL.Path = string(replacePath)
		}
		c.Next()
	}
}
