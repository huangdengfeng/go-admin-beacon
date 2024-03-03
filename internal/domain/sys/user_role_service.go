package sys

import (
	"context"
	"go-admin-beacon/internal/domain/dao"
	"go-admin-beacon/internal/infrastructure/constants"
)

// UserRoleService 用户角色服务
type UserRoleService struct {
	sysRoleDao *dao.SysRoleDao
}

var UserRoleServiceInstance = &UserRoleService{
	dao.SysRoleDaoInstance,
}

func (s *UserRoleService) FindRolesByUid(ctx context.Context, uid int32) ([]*dao.SysRolePO, error) {
	if constants.IsSuperAdmin(uid) {
		return s.sysRoleDao.FindValidRoles(ctx)
	}
	return s.sysRoleDao.FindRolesByUid(ctx, uid)
}
