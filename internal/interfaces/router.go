package interfaces

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-admin-beacon/internal/application/sys/auth"
	"go-admin-beacon/internal/infrastructure/config"
	"time"
)

func CreateRouter() *gin.Engine {
	if !config.DebugEnable {
		gin.SetMode(gin.ReleaseMode)
	}
	corsProperties := config.Global.Cors
	router := gin.Default()
	if corsProperties.Enable {
		router.Use(cors.New(cors.Config{
			AllowOrigins:     corsProperties.AllowOrigins,
			AllowCredentials: corsProperties.AllowCredentials,
			MaxAge:           corsProperties.MaxAge * time.Second,
			AllowWildcard:    true,
		}))
	}

	router.Use(auth.Filter)
	sysLoginApi := &sysLoginApi{}
	sysUserApi := &sysUserApi{}
	sysParamApi := &sysParamApi{}

	// 分组，也可以直接router.POST(xxxx)
	sys := router.Group("/sys")
	{
		sys.POST("/login/user_passwd", sysLoginApi.UserPasswdLogin)
		sys.POST("/user/page", sysUserApi.UserPageQry)
		sys.GET("/user/my", sysUserApi.My)
		sys.GET("/param/qry", sysParamApi.QryParam)
	}

	return router
}
