package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jiaruling/Gateway/controller"
	m "github.com/jiaruling/Gateway/middleware"
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

	oauth := router.Group("/oauth")
	oauth.Use(m.TranslationMiddleware())
	{
		controller.OAuthRegister(oauth)
	}

	router.Use(
		hm.HTTPAccessModeMiddleware(),     // done:匹配服务
		hm.HTTPFlowCountMiddleware(),      // done:流量统计
		hm.HTTPFlowLimitMiddleware(),      // done:服务限流
		hm.HTTPJwtAuthTokenMiddleware(),   // done:租户token检验
		hm.HTTPJwtFlowCountMiddleware(),   // done:租户流量统计
		hm.HTTPJwtFlowLimitMiddleware(),   // done:租户限流
		hm.HTTPWhiteListMiddleware(),      // done:白名单
		hm.HTTPBlackListMiddleware(),      // done:黑名单
		hm.HTTPHeaderTransferMiddleware(), // done:header头设置
		hm.HTTPStripUriMiddleware(),       // done:去除接入前缀
		hm.HTTPUrlRewriteMiddleware(),     // done:url重写
		hm.HTTPReverseProxyMiddleware(),   // tag:【重点】http反向代理
	)
	return router
}
