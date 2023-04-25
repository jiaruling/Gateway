package router

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	m "github.com/jiaruling/Gateway/middleware"
	"github.com/jiaruling/golang_utils/lib"
)

var (
	HttpSrvHandler  *http.Server
	HttpsSrvHandler *http.Server
)

func HttpServerRun() {
	gin.SetMode(lib.ViperConfMap["proxy"].GetString("base.debug_mode"))
	r := InitRouter(m.RecoveryMiddleware(), m.RequestLog())
	HttpSrvHandler = &http.Server{
		Addr:           lib.ViperConfMap["proxy"].GetString("http.addr"),
		Handler:        r,
		ReadTimeout:    time.Duration(lib.ViperConfMap["proxy"].GetInt("http.read_timeout")) * time.Second,
		WriteTimeout:   time.Duration(lib.ViperConfMap["proxy"].GetInt("http.write_timeout")) * time.Second,
		MaxHeaderBytes: 1 << uint(lib.ViperConfMap["proxy"].GetInt("http.max_header_bytes")),
	}
	log.Printf(" [INFO] http_proxy_run %s\n", lib.ViperConfMap["proxy"].GetString("http.addr"))
	if err := HttpSrvHandler.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf(" [ERROR] http_proxy_run %s err:%v\n", lib.ViperConfMap["proxy"].GetString("http.addr"), err)
	}
}

func HttpsServerRun() {
	gin.SetMode(lib.ViperConfMap["proxy"].GetString("base.debug_mode"))
	r := InitRouter(m.RecoveryMiddleware(), m.RequestLog())
	HttpsSrvHandler = &http.Server{
		Addr:           lib.ViperConfMap["proxy"].GetString("https.addr"),
		Handler:        r,
		ReadTimeout:    time.Duration(lib.ViperConfMap["proxy"].GetInt("https.read_timeout")) * time.Second,
		WriteTimeout:   time.Duration(lib.ViperConfMap["proxy"].GetInt("https.write_timeout")) * time.Second,
		MaxHeaderBytes: 1 << uint(lib.ViperConfMap["proxy"].GetInt("https.max_header_bytes")),
	}
	log.Printf(" [INFO] https_proxy_run %s\n", lib.ViperConfMap["proxy"].GetString("https.addr"))
	//todo 以下命令只在编译机有效，如果是交叉编译情况下需要单独设置路径
	//if err := HttpsSrvHandler.ListenAndServeTLS(cert_file.Path("server.crt"), cert_file.Path("server.key")); err != nil && err!=http.ErrServerClosed {
	if err := HttpsSrvHandler.ListenAndServeTLS("./cert_file/server.crt", "./cert_file/server.key"); err != nil && err != http.ErrServerClosed {
		log.Fatalf(" [ERROR] https_proxy_run %s err:%v\n", lib.ViperConfMap["proxy"].GetString("https.addr"), err)
	}
}

func HttpServerStop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := HttpSrvHandler.Shutdown(ctx); err != nil {
		log.Printf(" [ERROR] http_proxy_stop err:%v\n", err)
	}
	log.Printf(" [INFO] http_proxy_stop %v stopped\n", lib.ViperConfMap["proxy"].GetString("http.addr"))
}

func HttpsServerStop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := HttpsSrvHandler.Shutdown(ctx); err != nil {
		log.Fatalf(" [ERROR] https_proxy_stop err:%v\n", err)
	}
	log.Printf(" [INFO] https_proxy_stop %v stopped\n", lib.ViperConfMap["proxy"].GetString("https.addr"))
}
