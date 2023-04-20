package middleware

import (
	"errors"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jiaruling/Gateway/global"
)

func SessionAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if name, ok := session.Get(global.AdminSessionInfoKey).(string); !ok || name == "" {
			ResponseError(c, InternalErrorCode, errors.New("user not login"))
			c.Abort()
			return
		}
		c.Next()
	}
}
