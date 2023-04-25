package router

import (
	"github.com/gin-gonic/gin"
	hm "github.com/jiaruling/Gateway/proxy/http/middleware"
)

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	router := gin.New()
	router.Use(middlewares...)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Use(
		hm.HTTPAccessModeMiddleware(),
		// hm.HTTPFlowCountMiddleware(),
		// hm.HTTPFlowLimitMiddleware(),
		// hm.HTTPJwtAuthTokenMiddleware(),
		// hm.HTTPJwtFlowCountMiddleware(),
		// hm.HTTPJwtFlowLimitMiddleware(),
		// hm.HTTPWhiteListMiddleware(),
		// hm.HTTPBlackListMiddleware(),
		// hm.HTTPHeaderTransferMiddleware(),
		// hm.HTTPStripUriMiddleware(),
		// hm.HTTPUrlRewriteMiddleware(),
		hm.HTTPReverseProxyMiddleware(),
	)
	return router
}
