package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jiaruling/Gateway/global"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/override/docs"
)

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	// programatically set swagger info
	docs.SwaggerInfo.Title = global.ConfigBase.Swagger.Title
	docs.SwaggerInfo.Description = global.ConfigBase.Swagger.Desc
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = global.ConfigBase.Swagger.Host
	docs.SwaggerInfo.BasePath = global.ConfigBase.Swagger.BasePath
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router := gin.Default()
	router.Use(middlewares...)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	//demo
	// v1 := router.Group("/demo")
	// v1.Use()
	// {
	// 	// controller.DemoRegister(v1)
	// }

	// //非登陆接口
	// store := sessions.NewCookieStore([]byte("secret"))
	// apiNormalGroup := router.Group("/api")
	// apiNormalGroup.Use(sessions.Sessions("mysession", store),
	// 	// middleware.RecoveryMiddleware(),
	// 	// middleware.RequestLog(),
	// 	// middleware.TranslationMiddleware())
	// {
	// 	// controller.ApiRegister(apiNormalGroup)
	// }

	// //登陆接口
	// apiAuthGroup := router.Group("/api")
	// apiAuthGroup.Use(
	// 	sessions.Sessions("mysession", store),
	// 	// middleware.RecoveryMiddleware(),
	// 	// middleware.RequestLog(),
	// 	// middleware.SessionAuthMiddleware(),
	// 	// middleware.TranslationMiddleware()
	// )
	// {
	// 	// controller.ApiLoginRegister(apiAuthGroup)
	// }
	return router
}
