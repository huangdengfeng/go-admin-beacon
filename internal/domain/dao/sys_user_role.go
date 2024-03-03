package dao

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go-admin-beacon/internal/infrastructure/errors"
	"time"
)

type SysUserRole struct {
	Uid        int32 `gorm:"primaryKey"`
	RoleId     int32 `gorm:"primaryKey"`
	CreateUser int32
	CreateTime time.Time
}

func (s SysUserRole) TableName() string {
	return "sys_user_role"
}

type SysUserRoleDao struct {
}

var SysUserRoleDaoInstance = &SysUserRoleDao{}

// DeleteUserRoleByUid 根据uid 删除用户角色
func (s *SysUserRoleDao) DeleteUserRoleByUid(ctx context.Context, uid int32) error {
	result := getDbFromContext(ctx).Where("uid = ?", uid).Delete(&SysUserRole{})
	log.Infof("delete sys_user_role uid [%d] affect rows: %d", uid, result.RowsAffected)
	return result.Error
}

func (s *SysUserRoleDao) Save(ctx context.Context, pos []*SysUserRole) error {
	length := len(pos)
	if length <= 0 {
		return nil
	}
	result := getDbFromContext(ctx).Create(pos)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected != int64(length) {
		return errors.RowsAffectedNotMatch
	}
	return nil
}
