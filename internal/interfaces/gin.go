package interfaces

// 处理gin web 中间建逻辑
import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-admin-beacon/internal/application/sys/auth"
	"go-admin-beacon/internal/infrastructure/config"
	"go-admin-beacon/internal/infrastructure/errors"
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
	// 统一错误处理
	router.Use(func(c *gin.Context) {
		for _, e := range c.Errors {
			// 无权限，则转换为403
			if e.Err == errors.NoPermission {
				c.Status(http.StatusForbidden)
			} else if e.Err == errors.LoginSessionInvalid {
				c.Status(http.StatusUnauthorized)

			} else {
				c.JSON(http.StatusOK, response.Error(e.Err))
			}
			return
		}
		c.Next()
	})
	sysLoginApi := &sysLoginApi{}
	sysUserApi := &sysUserApi{}
	sysParamApi := &sysParamApi{}
	sysRoleApi := &sysRoleApi{}
	sysPermissionApi := &sysPermissionApi{}

	// 分组，也可以直接router.POST(xxxx)
	sys := router.Group("/sys")
	{
		sys.POST("/login/user_passwd", sysLoginApi.UserPasswdLogin)
		sys.POST("/user/page", sysUserApi.UserPageQry)
		sys.GET("/user/my", sysUserApi.My)
		sys.POST("/user/add", sysUserApi.AddUser)
		sys.POST("/user/modify", sysUserApi.ModifyUserCmdExe)
		sys.POST("/user/modify_pwd", sysUserApi.ModifyUserPwd)
		sys.GET("/user/detail/:uid", sysUserApi.UserDetailQry)

		sys.GET("/param/qry", sysParamApi.QryParam)

		sys.POST("/role/list", sysRoleApi.qryRoleList)
		sys.POST("/permission/list", sysPermissionApi.qryPermissionList)
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
