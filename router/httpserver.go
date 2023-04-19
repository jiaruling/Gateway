package router

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jiaruling/Gateway/global"
)

var (
	HttpSrvHandler *http.Server
)

func HttpServerRun() {
	if global.ConfigBase.Base.DebugMode == "debug" {
		gin.SetMode(gin.DebugMode)
		// 控制台显示日志显示颜色
		gin.ForceConsoleColor()
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	r := InitRouter()
	HttpSrvHandler = &http.Server{
		Addr:           global.ConfigBase.Http.Addr,
		Handler:        r,
		ReadTimeout:    time.Duration(global.ConfigBase.Http.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(global.ConfigBase.Http.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << uint(global.ConfigBase.Http.MaxHeaderBytes),
	}
	go func() {
		log.Printf(" [INFO] HttpServerRun:%s\n", global.ConfigBase.Http.Addr)
		if err := HttpSrvHandler.ListenAndServe(); err != nil {
			log.Fatalf(" [ERROR] HttpServerRun:%s err:%v\n", global.ConfigBase.Http.Addr, err)
		}
	}()
}

func HttpServerStop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := HttpSrvHandler.Shutdown(ctx); err != nil {
		log.Fatalf(" [ERROR] HttpServerStop err:%v\n", err)
	}
	log.Printf(" [INFO] HttpServerStop stopped\n")
}
