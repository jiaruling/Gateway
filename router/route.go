package router

import (
	"github.com/e421083458/gin_scaffold/controller"
	"github.com/e421083458/gin_scaffold/docs"
	"github.com/e421083458/gin_scaffold/middleware"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	// programatically set swagger info
	docs.SwaggerInfo.Title = lib.GetStringConf("base.swagger.title")
	docs.SwaggerInfo.Description = lib.GetStringConf("base.swagger.desc")
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = lib.GetStringConf("base.swagger.host")
	docs.SwaggerInfo.BasePath = lib.GetStringConf("base.swagger.base_path")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router := gin.Default()
	router.Use(middlewares...)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//demo
	v1 := router.Group("/demo")
	v1.Use(middleware.RecoveryMiddleware(), middleware.RequestLog(), middleware.IPAuthMiddleware(), middleware.TranslationMiddleware())
	{
		controller.DemoRegister(v1)
	}

	//非登陆接口
	store := sessions.NewCookieStore([]byte("secret"))
	apiNormalGroup := router.Group("/api")
	apiNormalGroup.Use(sessions.Sessions("mysession", store),
		middleware.RecoveryMiddleware(),
		middleware.RequestLog(),
		middleware.TranslationMiddleware())
	{
		controller.ApiRegister(apiNormalGroup)
	}

	//登陆接口
	apiAuthGroup := router.Group("/api")
	apiAuthGroup.Use(
		sessions.Sessions("mysession", store),
		middleware.RecoveryMiddleware(),
		middleware.RequestLog(),
		middleware.SessionAuthMiddleware(),
		middleware.TranslationMiddleware())
	{
		controller.ApiLoginRegister(apiAuthGroup)
	}
	return router
}
