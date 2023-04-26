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
		hm.HTTPAccessModeMiddleware(), // 匹配服务
		hm.HTTPFlowCountMiddleware(),  // 流量统计
		hm.HTTPFlowLimitMiddleware(),  // 服务限流
		// hm.HTTPJwtAuthTokenMiddleware(),
		// hm.HTTPJwtFlowCountMiddleware(),
		// hm.HTTPJwtFlowLimitMiddleware(),
		hm.HTTPWhiteListMiddleware(),      // 白名单
		hm.HTTPBlackListMiddleware(),      // 黑名单
		hm.HTTPHeaderTransferMiddleware(), // header头设置
		hm.HTTPStripUriMiddleware(),       // 去除接入前缀
		hm.HTTPUrlRewriteMiddleware(),     // url重写
		hm.HTTPReverseProxyMiddleware(),   // tag:【重点】http反向代理
	)
	return router
}
