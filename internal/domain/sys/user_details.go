package sys

// 用户信息
import (
	"context"
	"database/sql"
	"go-admin-beacon/internal/domain/dao"
	"go-admin-beacon/internal/infrastructure/constants"
	"go-admin-beacon/internal/infrastructure/errors"
)

type UserDetailsVO struct {
	UserName        string
	UserId          int32
	SecretKey       string
	PermissionCodes []string
	// ROLE_前缀
	RoleCodes []string
}

type UserDetailsService struct {
	sysUserDao       *dao.SysUserDao
	sysRoleDao       *dao.SysRoleDao
	sysPermissionDao *dao.SysPermissionDao
}

var UserDetailsServiceInstance = &UserDetailsService{
	sysUserDao:       dao.SysUserDaoInstance,
	sysRoleDao:       dao.SysRoleDaoInstance,
	sysPermissionDao: dao.SysPermissionDaoInstance,
}

func (s *UserDetailsService) GetUserDetails(uid int32) (*UserDetailsVO, error) {
	var po *dao.SysUserPO
	var roles []*dao.SysRolePO
	var permissions []*dao.SysPermissionPO
	txFunc := func(ctx context.Context) error {
		var err error
		po, err = s.sysUserDao.FindByUid(ctx, uid)
		if nil != err {
			return err
		}
		// 用户不存在
		if po == nil {
			return errors.UserNotExists
		}
		roles, err = s.sysRoleDao.FindRolesByUid(ctx, po.Uid)
		if nil != err {
			return err
		}
		permissions, err = s.sysPermissionDao.FindPermissionsByUid(ctx, uid)
		if nil != err {
			return err
		}
		return nil
	}
	err := dao.DoTransaction(context.Background(), txFunc, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return nil, err
	}
	var roleCodes = make([]string, 0)
	var permissionCodes = make([]string, 0)
	for _, role := range roles {
		// 兼容java 写法，添加前缀
		roleCodes = append(roleCodes, constants.RolePrefix+role.Code)
	}

	for _, permission := range permissions {
		permissionCodes = append(permissionCodes, permission.Code)
	}

	// 用户角色
	details := &UserDetailsVO{
		UserName:        po.UserName,
		UserId:          po.Uid,
		SecretKey:       po.SecretKey,
		RoleCodes:       roleCodes,
		PermissionCodes: permissionCodes,
	}
	return details, nil
}
