package sys

import (
	"context"
	"go-admin-beacon/internal/application/sys/auth"
	"go-admin-beacon/internal/domain/sys"
	"go-admin-beacon/internal/infrastructure/response"
)

type ModifyUserPwdCmd struct {
	// Uid
	Uid int32 `json:"uid"`
	// Password 密码
	Password string `json:"password" binding:"required,min=6"`
}

type modifyUserPwdCmdExe struct {
	userService *sys.UserService
}

func NewModifyUserPwdCmdExe() *modifyUserPwdCmdExe {
	return &modifyUserPwdCmdExe{
		sys.UserServiceInstance,
	}
}

func (e *modifyUserPwdCmdExe) Execute(ctx context.Context, cmd *ModifyUserPwdCmd) (*response.Response, error) {
	if err := auth.CheckPermission(ctx, "sys:user:modify_pwd"); err != nil {
		return nil, err
	}
	userDetailsVO := auth.GetUserFromContext(ctx)
	if err := e.userService.ModifyUserPwd(ctx, cmd.Uid, cmd.Password, userDetailsVO.UserId); err != nil {
		return nil, err
	}
	return response.DefaultSuccess, nil
}
