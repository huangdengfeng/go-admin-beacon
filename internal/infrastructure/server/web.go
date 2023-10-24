package server

import (
	"context"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go-admin-beacon/internal/infrastructure/config"
	"net/http"
	"time"
)

// server的配置
var serverConf = &config.Global.Server

// Start 启动服务
func Start(router *gin.Engine) *http.Server {
	server := &http.Server{
		Addr:    serverConf.Listen,
		Handler: router,
	}
	go func() {
		// 启动&监听
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	return server
}

// Shutdown 优雅关闭服务
func Shutdown(server *http.Server) error {
	ctx, cancel := context.WithTimeout(context.Background(), serverConf.ShutdownTimeout*time.Second)
	defer cancel()
	err := server.Shutdown(ctx)
	return err
}
