package router

import (
	"fmt"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/jiaruling/Gateway/controller"
	"github.com/jiaruling/Gateway/docs"
	"github.com/jiaruling/Gateway/global"
	"github.com/jiaruling/Gateway/middleware"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	// done: 管理员登录登出
	adminLoginRouter := router.Group("/admin_login")
	store, err := redis.NewStore(10, "tcp", fmt.Sprintf("%s:%d", global.ConfigRedis.List["default"].Ip, global.ConfigRedis.List["default"].Port), global.ConfigRedis.List["default"].Password, []byte("secret"))
	if err != nil {
		log.Fatalf("sessions.NewRedisStore err:%v", err)
	}
	adminLoginRouter.Use(
		sessions.Sessions("mysession", store),
		middleware.RecoveryMiddleware(),
		middleware.RequestLog(),
		middleware.TranslationMiddleware())
	{
		controller.AdminLoginRegister(adminLoginRouter)
	}

	// done: 管理员信息及密码修改
	adminRouter := router.Group("/admin")
	adminRouter.Use(
		sessions.Sessions("mysession", store),
		middleware.RecoveryMiddleware(),
		middleware.RequestLog(),
		middleware.SessionAuthMiddleware(),
		middleware.TranslationMiddleware())
	{
		controller.AdminRegister(adminRouter)
	}

	// done: 服务
	serviceRouter := router.Group("/service")
	serviceRouter.Use(
		sessions.Sessions("mysession", store),
		middleware.RecoveryMiddleware(),
		middleware.RequestLog(),
		middleware.SessionAuthMiddleware(),
		middleware.TranslationMiddleware())
	{
		controller.ServiceRegister(serviceRouter)
	}

	// done: 租户
	appRouter := router.Group("/app")
	appRouter.Use(
		sessions.Sessions("mysession", store),
		middleware.RecoveryMiddleware(),
		middleware.RequestLog(),
		middleware.SessionAuthMiddleware(),
		middleware.TranslationMiddleware())
	{
		controller.APPRegister(appRouter)
	}

	// done: 首页大盘
	dashRouter := router.Group("/dashboard")
	dashRouter.Use(
		sessions.Sessions("mysession", store),
		middleware.RecoveryMiddleware(),
		middleware.RequestLog(),
		middleware.SessionAuthMiddleware(),
		middleware.TranslationMiddleware())
	{
		controller.DashboardRegister(dashRouter)
	}

	return router
}
