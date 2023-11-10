package interfaces

import (
	"github.com/gin-gonic/gin"
	"go-admin-beacon/internal/application/sys/auth"
	"go-admin-beacon/internal/infrastructure/config"
)

func CreateRouter() *gin.Engine {
	if !config.DebugEnable {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	router.Use(auth.Filter)
	sysLoginApi := &sysLoginApi{}
	sysUserApi := &sysUserApi{}

	// 分组，也可以直接router.POST(xxxx)
	sys := router.Group("/sys")
	{
		sys.POST("/login/user_passwd", sysLoginApi.UserPasswdLogin)
		sys.POST("/user/page", sysUserApi.UserPageQry)
		sys.GET("/user/my", sysUserApi.My)
	}

	return router
}
