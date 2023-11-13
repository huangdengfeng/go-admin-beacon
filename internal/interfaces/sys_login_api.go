package interfaces

import (
	"github.com/gin-gonic/gin"
	"go-admin-beacon/internal/application/sys"
	"go-admin-beacon/internal/infrastructure/errors"
	"go-admin-beacon/internal/infrastructure/response"
	"net/http"
)

type sysLoginApi struct {
}

var userPasswdLoginExe = sys.NewUserPasswdLoginExe()

// UserPasswdLogin 账号密码登录
func (s *sysLoginApi) UserPasswdLogin(c *gin.Context) {
	cmd := &sys.UserPasswdLoginCmd{}
	if err := c.ShouldBindJSON(cmd); err != nil {
		c.JSON(http.StatusOK, response.ErrorWithCodeMsg(errors.BadArgs.Code, err.Error()))
		return
	}
	response, err := userPasswdLoginExe.Execute(cmd)
	packResponse(c, response, err)
}
