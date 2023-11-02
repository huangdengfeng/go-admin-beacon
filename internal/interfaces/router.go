package interfaces

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go-admin-beacon/internal/infrastructure/config"
	"go-admin-beacon/internal/infrastructure/security"
	"net/http"
	"strings"
)

func CreateRouter() *gin.Engine {
	if !config.DebugEnable {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	router.Use(auth)
	sysLoginApi := &sysLoginApi{}
	sysUserApi := &sysUserApi{}

	// 分组，也可以直接router.POST(xxxx)
	sys := router.Group("/sys")
	{
		sys.POST("/login/user_passwd", sysLoginApi.UserPasswdLogin)
		sys.POST("/user/page", sysUserApi.UserPageQry)
	}

	return router
}

var unAuthUris = [...]string{"/public/", "/sys/login/"}

func auth(c *gin.Context) {
	path := c.Request.URL.Path
	var checkAuth = true
	for _, uri := range unAuthUris {
		if strings.HasPrefix(path, uri) {
			checkAuth = false
			break
		}
	}
	// url 需要验证登录
	if !checkAuth {
		c.Next()
		return
	}
	// Authorization: Bearer xxxxx
	authorizationHeader := c.Request.Header.Get("Authorization")
	if authorizationHeader == "" {
		c.Status(http.StatusUnauthorized)
		c.Abort()
		return
	}
	token := strings.TrimPrefix(authorizationHeader, "Bearer ")

	tokenInfo, err := security.ParseToken(token, []byte(config.Global.Login.TokenSignKey))
	if err != nil {
		log.Errorf("parse security token error %s ", err.Error())
		c.Status(http.StatusUnauthorized)
		c.Abort()
		return
	}
	log.Debugf("parse security token success:%+v", tokenInfo)
}
