package sys

// 账号密码登录

import (
	"context"
	"go-admin-beacon/internal/domain/sys"
	"go-admin-beacon/internal/infrastructure/errors"
	"go-admin-beacon/internal/infrastructure/response"
)

type UserPasswdLoginCmd struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6,max=50"`
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

func (e *userPasswdLoginExe) Execute(ctx context.Context, cmd *UserPasswdLoginCmd) (*response.Response, error) {
	verify, err := e.userPasswdVerifyService.Verify(ctx, cmd.Username, cmd.Password)
	if nil != err {
		return nil, err
	}
	// 用户名密码错误
	if !verify {
		return response.Error(errors.UserPasswdWrong), nil
	}
	// 颁发token
	token, err := e.userTokenService.Create(ctx, cmd.Username)
	if nil != err {
		return nil, err
	}
	return response.Success(&AuthorizationTokenCO{token}), nil
}
