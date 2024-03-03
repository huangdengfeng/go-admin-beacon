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
	sysUserDao     *dao.SysUserDao
	sysUserRoleDao *dao.SysUserRoleDao
}

var UserServiceInstance = &UserService{
	sysUserDao:     dao.SysUserDaoInstance,
	sysUserRoleDao: dao.SysUserRoleDaoInstance,
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
	RoleIds  []int32
}

// ModifyUserVO 修改用户
type ModifyUserVO struct {
	Uid      int32
	UserName string
	Name     string
	Mobile   string
	Email    string
	Photo    string
	Status   int8
	Remark   string
	RoleIds  []int32
}

func (s *UserService) AddUser(ctx context.Context, vo *AddUserVO, operator int32) (int32, error) {
	var uid int32
	txFunc := func(ctx context.Context) error {
		user, err := s.sysUserDao.FindByUserName(ctx, vo.UserName)
		if err != nil {
			return err
		}
		if user != nil {
			return errors.UserNameExists
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
		return s.SaveUserRoles(ctx, uid, vo.RoleIds, operator)
	}
	return uid, dao.DoTransaction(ctx, txFunc)
}

func (s *UserService) SaveUserRoles(ctx context.Context, uid int32, roleIds []int32, operator int32) error {
	if len(roleIds) > 0 {
		userRoles := make([]*dao.SysUserRole, 0, len(roleIds))
		for _, roleId := range roleIds {
			userRoles = append(userRoles, &dao.SysUserRole{
				Uid:        uid,
				RoleId:     roleId,
				CreateUser: operator,
				CreateTime: time.Now(),
			})
		}

		return dao.DoTransaction(ctx, func(ctx context.Context) error {
			return s.sysUserRoleDao.Save(ctx, userRoles)
		})
	}
	return nil
}

// ModifyUserPwd 修改指定用户密码
func (s *UserService) ModifyUserPwd(ctx context.Context, uid int32, password string, operator int32) error {
	encodedPassword, err := security.EncodePassword(password)
	if err != nil {
		return err
	}
	txFun := func(ctx context.Context) error {
		po, err := s.sysUserDao.FindByUid(ctx, uid)
		if err != nil {
			return err
		}
		if po == nil {
			return errors.UserNotExists
		}
		po.Password = encodedPassword
		po.UpdateTime = time.Now()
		po.UpdateUser = operator
		err = s.sysUserDao.Update(ctx, po)
		return err
	}
	return dao.DoTransaction(ctx, txFun)
}

func (s *UserService) ModifyUser(ctx context.Context, vo *ModifyUserVO, operator int32) error {
	uid := vo.Uid
	txFun := func(ctx context.Context) error {
		userPO, err := s.sysUserDao.FindByUid(ctx, uid)
		if err != nil {
			return err
		}
		if userPO == nil {
			return errors.UserNotExists
		}
		existsUserPO, err := s.sysUserDao.FindByUserName(ctx, vo.UserName)
		if err != nil {
			return err
		}
		// 用户名存在且不是自己
		if existsUserPO != nil && existsUserPO.Uid != uid {
			return errors.UserNameExists
		}
		// RoleIds  []string
		userPO.UserName = vo.UserName
		userPO.Name = vo.Name
		userPO.Mobile = vo.Mobile
		userPO.Email = vo.Email
		userPO.Photo = vo.Photo
		userPO.Status = vo.Status
		userPO.Remark = vo.Remark
		userPO.UpdateUser = operator
		userPO.UpdateTime = time.Now()
		err = s.sysUserDao.Update(ctx, userPO)
		if err != nil {
			return err
		}
		// 处理角色
		err = s.sysUserRoleDao.DeleteUserRoleByUid(ctx, userPO.Uid)
		if err != nil {
			return err
		}
		return s.SaveUserRoles(ctx, uid, vo.RoleIds, operator)
	}
	return dao.DoTransaction(ctx, txFun)
}
