package sys

import (
	"go-admin-beacon/internal/domain/dao"
)

type UserPasswdVerifyService struct {
	sysUserDao *dao.SysUserDao
}

var UserPasswdVerifyServiceInstance = &UserPasswdVerifyService{
	sysUserDao: dao.SysUserDaoInstance,
}

func (s *UserPasswdVerifyService) Verify(userName string, password string) (bool, error) {
	if po, err := s.sysUserDao.FindByUserName(userName); nil != err {
		return false, err
	} else {
		// 用户不存在
		if nil == po {
			return false, nil
		}
		// 验证密码
		return true, nil
	}
}
