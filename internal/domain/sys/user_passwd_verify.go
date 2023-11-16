package sys

import (
	"context"
	"go-admin-beacon/internal/domain/dao"
	"go-admin-beacon/internal/infrastructure/security"
)

type UserPasswdVerifyService struct {
	sysUserDao *dao.SysUserDao
}

var UserPasswdVerifyServiceInstance = &UserPasswdVerifyService{
	sysUserDao: dao.SysUserDaoInstance,
}

func (s *UserPasswdVerifyService) Verify(ctx context.Context, userName string, password string) (bool, error) {
	po, err := s.sysUserDao.FindByUserName(ctx, userName)
	if nil != err {
		return false, err
	}
	// 用户不存在
	if po == nil {
		return false, nil
	}
	return security.PasswordMatches(password, po.Password), nil
}
