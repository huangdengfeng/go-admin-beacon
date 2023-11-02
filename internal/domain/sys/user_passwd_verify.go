package sys

import (
	"go-admin-beacon/internal/domain/dao"
	"go-admin-beacon/internal/infrastructure/security"
)

type UserPasswdVerifyService struct {
	sysUserDao *dao.SysUserDao
}

var UserPasswdVerifyServiceInstance = &UserPasswdVerifyService{
	sysUserDao: dao.SysUserDaoInstance,
}

func (s *UserPasswdVerifyService) Verify(userName string, password string) (bool, error) {
	po, err := s.sysUserDao.FindByUserName(userName)
	if nil != err {
		return false, err
	}
	// 用户不存在
	if po == nil {
		return false, nil
	}
	return security.PasswordMatches(password, po.Password), nil
}
