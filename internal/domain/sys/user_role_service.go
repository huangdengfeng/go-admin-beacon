package sys

import (
	"context"
	"go-admin-beacon/internal/domain/dao"
	"go-admin-beacon/internal/infrastructure/constants"
)

// UserPermissionService 用户角色服务
type UserPermissionService struct {
	sysRoleDao       *dao.SysRoleDao
	sysPermissionDao *dao.SysPermissionDao
}

var UserPermissionServiceInstance = &UserPermissionService{
	sysRoleDao:       dao.SysRoleDaoInstance,
	sysPermissionDao: dao.SysPermissionDaoInstance,
}

func (s *UserPermissionService) FindRolesByUid(ctx context.Context, uid int32) ([]*dao.SysRolePO, error) {
	if constants.IsSuperAdmin(uid) {
		return s.sysRoleDao.FindValidRoles(ctx)
	}
	return s.sysRoleDao.FindRolesByUid(ctx, uid)
}

func (s *UserPermissionService) FindPermissionsByUid(ctx context.Context, uid int32) ([]*dao.SysPermissionPO, error) {
	if constants.IsSuperAdmin(uid) {
		return s.sysPermissionDao.FindValidPermissions(ctx)
	} else {
		return s.sysPermissionDao.FindPermissionsByUid(ctx, uid)
	}
}
