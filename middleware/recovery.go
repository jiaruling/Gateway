package middleware

import (
	"errors"
	"fmt"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/jiaruling/Gateway/global"
	"github.com/jiaruling/Gateway/public"
)

// RecoveryMiddleware捕获所有panic，并且返回错误信息
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				//先做一下日志记录
				fmt.Println(string(debug.Stack()))
				trace, log := public.GetTraceAndLog()
				log.Info(trace, "_com_panic", map[string]interface{}{
					"error": fmt.Sprint(err),
					"stack": string(debug.Stack()),
				})

				if global.ConfigBase.Base.DebugMode != "debug" {
					ResponseError(c, 500, errors.New("内部错误"))
					return
				} else {
					ResponseError(c, 500, errors.New(fmt.Sprint(err)))
					return
				}
			}
		}()
		c.Next()
	}
}
