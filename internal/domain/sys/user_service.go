package sys

import (
	"context"
	"go-admin-beacon/internal/domain/dao"
	"go-admin-beacon/internal/infrastructure/constants"
	"go-admin-beacon/internal/infrastructure/errors"
	"go-admin-beacon/internal/infrastructure/security"
	"time"
)

type UserService struct {
	sysUserDao *dao.SysUserDao
}

var UserServiceInstance = &UserService{
	sysUserDao: dao.SysUserDaoInstance,
}

// AddUserVO 添加用户值对象
type AddUserVO struct {
	UserName string
	Password string
	Name     string
	Mobile   string
	Email    string
	Photo    string
	Remark   string
	RoleIds  []string
}

func (s *UserService) AddUser(ctx context.Context, vo *AddUserVO, operator int32) (int32, error) {
	var uid int32
	txFunc := func(ctx context.Context) error {
		user, err := s.sysUserDao.FindByUserName(ctx, vo.UserName)
		if err != nil {
			return err
		}
		if user != nil {
			return errors.UserExists
		}
		var password string
		if vo.Password != "" {
			encodePassword, err := security.EncodePassword(vo.Password)
			if err != nil {
				return err
			}
			password = encodePassword
		}
		// 对象转换
		po := &dao.SysUserPO{
			UserName:   vo.UserName,
			Password:   password,
			SecretKey:  security.GenerateUserKey(),
			Name:       vo.Name,
			Mobile:     vo.Mobile,
			Email:      vo.Email,
			Photo:      vo.Photo,
			Status:     constants.DbNormalStatus,
			CreateUser: operator,
			CreateTime: time.Now(),
			UpdateUser: operator,
			UpdateTime: time.Now(),
			Remark:     vo.Remark,
		}
		userId, err := s.sysUserDao.Save(ctx, po)
		if err != nil {
			return err
		}
		uid = userId
		return nil
	}
	return uid, dao.DoTransaction(ctx, txFunc)
}
