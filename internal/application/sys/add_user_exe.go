package sys

// 用户分页信息查询

import (
	"context"
	"go-admin-beacon/internal/application/sys/auth"
	"go-admin-beacon/internal/domain/dao"
	"go-admin-beacon/internal/domain/sys"
	"go-admin-beacon/internal/infrastructure/errors"
	"go-admin-beacon/internal/infrastructure/response"
)

type AddUserCmd struct {
	UserName string   `json:"userName" binding:"required,max=50"`
	Password string   `json:"password" binding:"required,max=50"`
	Name     string   `json:"name" binding:"required,max=50"`
	Mobile   string   `json:"mobile" binding:"max=50"`
	Email    string   `json:"email"`
	Photo    string   `json:"photo"`
	Remark   string   `json:"remark"`
	RoleIds  []string `json:"roleIds"`
}

type addUserCmdExe struct {
	userService *sys.UserService
}

func NewAddUserCmdExe() *addUserCmdExe {
	return &addUserCmdExe{sys.UserServiceInstance}
}

func (e *addUserCmdExe) Execute(ctx context.Context, cmd *AddUserCmd) (*response.Response, error) {
	if !auth.CheckPermission(ctx, "sys:user:add") {
		return nil, errors.NoPermission
	}
	detailsVO := auth.GetUserFromContext(ctx)
	txFunc := func(ctx context.Context) error {
		_, err := e.userService.AddUser(ctx, &sys.AddUserVO{
			UserName: cmd.UserName,
			Password: cmd.Password,
			Name:     cmd.Name,
			Mobile:   cmd.Mobile,
			Email:    cmd.Email,
			Photo:    cmd.Photo,
			Remark:   cmd.Remark,
			RoleIds:  cmd.RoleIds,
		}, detailsVO.UserId)
		if err != nil {
			return err
		}
		return nil
	}
	err := dao.DoTransaction(ctx, txFunc)
	if err != nil {
		return nil, err
	}
	return response.DefaultSuccess, nil
}
