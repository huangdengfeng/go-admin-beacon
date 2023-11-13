package sys

import (
	"context"
	"go-admin-beacon/internal/application/sys/auth"
	"go-admin-beacon/internal/domain/dao"
	"go-admin-beacon/internal/infrastructure/response"
)

type MyInfoCO struct {
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
	// 角色编码
	Roles []string `json:"roles"`
	// 角色名称
	RoleNames []string `json:"roleNames"`
	// 权限编码
	Permissions []string `json:"permissions"`
}

type myInfoExe struct {
	sysUserDao *dao.SysUserDao
	sysRoleDao *dao.SysRoleDao
}

func NewMyInfoExe() *myInfoExe {
	return &myInfoExe{
		sysUserDao: dao.SysUserDaoInstance,
		sysRoleDao: dao.SysRoleDaoInstance,
	}
}

func (e *myInfoExe) Execute(ctx context.Context) (*response.Response, error) {
	userDetailsVO := auth.GetUserFromContext(ctx)
	po, err := e.sysUserDao.FindByUid(ctx, userDetailsVO.UserId)
	if err != nil {
		return nil, err
	}
	roles, err := e.sysRoleDao.FindRolesByUid(ctx, po.Uid)
	if err != nil {
		return nil, err
	}
	// var RoleNames []string 为nil ,json 字段为null 对前端不友好
	var RoleNames = make([]string, 0)
	for _, role := range roles {
		RoleNames = append(RoleNames, role.Name)
	}
	co := &MyInfoCO{
		UserName:    po.UserName,
		Name:        po.Name,
		Mobile:      po.Mobile,
		Email:       po.Email,
		Photo:       po.Photo,
		Roles:       userDetailsVO.RoleCodes,
		RoleNames:   RoleNames,
		Permissions: userDetailsVO.PermissionCodes,
	}
	return response.Success(co), nil

}
