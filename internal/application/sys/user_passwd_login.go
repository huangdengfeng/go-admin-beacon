package sys

// 账号密码登录

import (
	"go-admin-beacon/internal/domain/sys"
	"go-admin-beacon/internal/infrastructure/errors"
	"go-admin-beacon/internal/infrastructure/response"
)

type UserPasswdLoginCmd struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthorizationTokenCO struct {
	Token string `json:"token"`
}

type userPasswdLoginExe struct {
	userPasswdVerifyService *sys.UserPasswdVerifyService
	userTokenService        *sys.UserTokenService
}

func NewUserPasswdLoginExe() *userPasswdLoginExe {
	return &userPasswdLoginExe{userPasswdVerifyService: sys.UserPasswdVerifyServiceInstance, userTokenService: sys.UserTokenServiceInstance}
}

func (e *userPasswdLoginExe) Execute(cmd *UserPasswdLoginCmd) *response.Response {
	verify, err := e.userPasswdVerifyService.Verify(cmd.Username, cmd.Password)
	if nil != err {
		return response.Error(err)
	}
	// 用户名密码错误
	if !verify {
		return response.Error(errors.UserPasswdWrong)
	}
	// 颁发token
	token, err := e.userTokenService.Create(cmd.Username)
	if nil != err {
		return response.Error(err)
	}
	return response.Success(&AuthorizationTokenCO{token})
}
