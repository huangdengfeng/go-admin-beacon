package main

import (
	log "github.com/sirupsen/logrus"
	"go-admin-beacon/internal/infrastructure/config"
	"go-admin-beacon/internal/infrastructure/server"
	"go-admin-beacon/internal/interfaces"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//初始化自定义配置
	config.Init()
	defer config.Shutdown()
	// 监听
	httpServer := server.Start(interfaces.CreateRouter())
	// 常驻进程需要优雅退出
	// 监听正常退出信号
	quit := make(chan os.Signal, 1)
	// ctrl+c || kill
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Info("Application Stopping...")
	// 优雅关闭
	if err := server.Shutdown(httpServer); err != nil {
		log.Fatalf("Server Shutdown: %s", err)
	}
	log.Info("Application Stopped!!!")
}
