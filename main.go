package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/jiaruling/Gateway/router"
)

func main() {
	router.HttpServerRun()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	router.HttpServerStop()
}
