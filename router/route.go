package router

import (
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/jiaruling/Gateway/controller"
	"github.com/jiaruling/Gateway/global"
	"github.com/jiaruling/Gateway/middleware"
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

	adminLoginRouter := router.Group("/admin_login")
	store, err := redis.NewStore(10, "tcp", global.ConfigRedis.List["default"].ProxyList[0], global.ConfigRedis.List["default"].Password, []byte("secret"))
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

	return router
}
