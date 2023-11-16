package interfaces

// 处理gin web 中间建逻辑
import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-admin-beacon/internal/application/sys/auth"
	"go-admin-beacon/internal/infrastructure/config"
	"go-admin-beacon/internal/infrastructure/response"
	"net/http"
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
	sysRoleApi := &sysRoleApi{}

	// 分组，也可以直接router.POST(xxxx)
	sys := router.Group("/sys")
	{
		sys.POST("/login/user_passwd", sysLoginApi.UserPasswdLogin)
		sys.POST("/user/page", sysUserApi.UserPageQry)
		sys.GET("/user/my", sysUserApi.My)
		sys.POST("/user/add", sysUserApi.AddUser)

		sys.GET("/param/qry", sysParamApi.QryParam)

		sys.POST("/role/list", sysRoleApi.qryRoleList)
	}

	return router
}

func packResponse(c *gin.Context, resp *response.Response, err error) {
	if err == nil {
		c.JSON(http.StatusOK, resp)
		return
	}
	// gin 上下文放入错误，全局处理
	_ = c.Error(err)
}
