package interfaces

import (
	"github.com/gin-gonic/gin"
	"go-admin-beacon/internal/infrastructure/config"
)

func CreateRouter() *gin.Engine {
	if !config.DebugEnable {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	sysLoginApi := &sysLoginApi{}
	router.POST("/sys/login/user_passwd", sysLoginApi.UserPasswdLogin)
	return router
}
