package sys

import (
	"context"
	"go-admin-beacon/internal/application/sys/auth"
	"go-admin-beacon/internal/domain/dao"
	"go-admin-beacon/internal/domain/sys"
	"go-admin-beacon/internal/infrastructure/constants"
	"go-admin-beacon/internal/infrastructure/errors"
	"go-admin-beacon/internal/infrastructure/response"
	"time"
)

type UserDetailQry struct {
	Uid int32 `json:"uid" binding:"required"`
}

type UserDetailCO struct {
	// 用户ID
	Uid int32 `json:"uid"`
	// 用户名
	UserName string `json:"userName"`
	// 姓名
	Name string `json:"name"`
	// 手机号
	Mobile string `json:"mobile"`
	// 邮件
	Email string `json:"email"`
	// 照片
	Photo string `json:"photo"`
	// 状态
	Status int8 `json:"status"`
	// 状态名称
	StatusName string `json:"statusName"`
	// 创建时间 yyyy-MM-dd HH:mm:ss
	CreateTime string `json:"createTime"`
	// 修改时间 yyyy-MM-dd HH:mm:ss
	UpdateTime string `json:"updateTime"`
	// 备注
	Remarks string `json:"remarks"`
	// 角色
	RoleIds []int32 `json:"roleIds"`
}

type userDetailQryExe struct {
	sysUserDao      *dao.SysUserDao
	userRoleService *sys.UserPermissionService
}

func NewUserDetailQryExe() *userDetailQryExe {
	return &userDetailQryExe{
		sysUserDao:      dao.SysUserDaoInstance,
		userRoleService: sys.UserPermissionServiceInstance,
	}
}

func (e *userDetailQryExe) Execute(ctx context.Context, qry *UserDetailQry) (*response.Response, error) {
	if err := auth.CheckPermission(ctx, "sys:user:qry"); err != nil {
		return nil, err
	}
	var userPO *dao.SysUserPO
	var sysRolePOs []*dao.SysRolePO
	txFunc := func(ctx context.Context) error {
		var err error
		// 查询用户信息
		userPO, err = e.sysUserDao.FindByUid(ctx, qry.Uid)
		if err != nil {
			return err
		}
		if userPO == nil {
			return errors.UserNotExists
		}
		// 查询角色列表
		sysRolePOs, err = e.userRoleService.FindRolesByUid(ctx, userPO.Uid)
		if err != nil {
			return err
		}
		return nil
	}
	if err := dao.DoTransaction(ctx, txFunc); err != nil {
		return nil, err
	}
	var roleIds = make([]int32, 0)
	for _, role := range sysRolePOs {
		roleIds = append(roleIds, role.Id)
	}
	co := &UserDetailCO{
		Uid:        userPO.Uid,
		UserName:   userPO.UserName,
		Name:       userPO.Name,
		Mobile:     userPO.Mobile,
		Email:      userPO.Email,
		Photo:      userPO.Photo,
		Status:     userPO.Status,
		StatusName: constants.GetDictName(constants.DictSysUserStatus, userPO.Status),
		CreateTime: userPO.CreateTime.Format(time.DateTime),
		UpdateTime: userPO.UpdateTime.Format(time.DateTime),
		Remarks:    userPO.Remark,
		RoleIds:    roleIds,
	}

	return response.Success(co), nil

}
