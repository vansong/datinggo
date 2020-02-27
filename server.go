package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"goinspect.cn/app/conf"
	"goinspect.cn/app/http/route"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	gin.SetMode(conf.Conf.AppMode)
	engine := gin.New()

	// reg http routes
	route.SetupRouter(engine)

	// init & start http server
	server := &http.Server{
		Addr:              conf.Conf.AppPort,
		Handler:           engine,
	}
	fmt.Println("Server Start Successully")
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP Server listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅关闭服务器（设置5秒的超时时间）
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, os.Interrupt)
	sig := <-signalChan
	log.Println("Get Signal:", sig)
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server Exiting")
}