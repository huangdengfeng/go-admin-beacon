package sys

import (
	"context"
	"go-admin-beacon/internal/application/sys/auth"
	"go-admin-beacon/internal/domain/sys"
	"go-admin-beacon/internal/infrastructure/errors"
	"go-admin-beacon/internal/infrastructure/response"
)

type ModifyUserCmd struct {
	Uid      int32   `json:"uid"`
	UserName string  `json:"userName" binding:"required,max=50"`
	Name     string  `json:"name" binding:"required,max=50"`
	Mobile   string  `json:"mobile" binding:"max=50"`
	Email    string  `json:"email"`
	Photo    string  `json:"photo"`
	Status   int8    `json:"status"`
	Remark   string  `json:"remark"`
	RoleIds  []int32 `json:"roleIds"`
}
type modifyUserCmdExe struct {
	userService *sys.UserService
}

func NewModifyUserCmdExe() *modifyUserCmdExe {
	return &modifyUserCmdExe{
		sys.UserServiceInstance,
	}
}

func (e *modifyUserCmdExe) Execute(ctx context.Context, cmd *ModifyUserCmd) (*response.Response, error) {
	if !auth.CheckPermission(ctx, "sys:user:modify") {
		return nil, errors.NoPermission
	}
	userDetailsVO := auth.GetUserFromContext(ctx)
	if err := e.userService.ModifyUser(ctx,
		&sys.ModifyUserVO{
			Uid:      cmd.Uid,
			UserName: cmd.UserName,
			Name:     cmd.Name,
			Mobile:   cmd.Mobile,
			Email:    cmd.Email,
			Photo:    cmd.Photo,
			Status:   cmd.Status,
			Remark:   cmd.Remark,
			RoleIds:  cmd.RoleIds,
		}, userDetailsVO.UserId); err != nil {
		return nil, err
	}
	return response.DefaultSuccess, nil
}
