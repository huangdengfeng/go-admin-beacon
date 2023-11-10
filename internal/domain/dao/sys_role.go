package dao

import (
	"context"
	"go-admin-beacon/internal/infrastructure/constants"
	"time"
)

type SysRolePO struct {
	Id   int32 `gorm:"primaryKey"`
	Code string
	Name string
	// Not null
	Status     int8
	CreateUser int32
	CreateTime time.Time
	UpdateUser int32
	UpdateTime time.Time
}

// TableName 自定义表名
func (s SysRolePO) TableName() string {
	return "sys_role"
}

type SysRoleDao struct {
	*dao
}

var SysRoleDaoInstance = &SysRoleDao{&dao{db: getDb}}

func (s *SysRoleDao) FindRolesByUid(context context.Context, uid int32) ([]*SysRolePO, error) {
	var roles []*SysRolePO
	result := getDbFromContext(context).Raw("SELECT * FROM sys_role sr WHERE sr.status = ? AND "+
		"EXISTS(SELECT 1 FROM sys_user_role sur WHERE sur.role_id = sr.id AND sur.uid = ?)", constants.DbNormalStatus, uid).Scan(&roles)
	if result.Error != nil {
		return nil, result.Error
	}
	return roles, nil
}
