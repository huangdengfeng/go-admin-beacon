package dao

import (
	"context"
	"go-admin-beacon/internal/infrastructure/constants"
	"time"
)

type SysPermissionPO struct {
	Id   int32 `gorm:"primaryKey"`
	Code string
	Name string
	// Null for root
	ParentId  *int32
	ParentIds string
	// Not null
	Status     int8
	CreateUser int32
	CreateTime time.Time
	UpdateUser int32
	UpdateTime time.Time
}

// TableName 自定义表名
func (s SysPermissionPO) TableName() string {
	return "sys_permission"
}

type SysPermissionDao struct {
}

var SysPermissionDaoInstance = &SysPermissionDao{}

func (s *SysPermissionDao) FindPermissionsByUid(context context.Context, uid int32) ([]*SysPermissionPO, error) {
	var permissions []*SysPermissionPO
	sql := "SELECT * FROM sys_permission sp WHERE sp.status = ? AND  EXISTS(SELECT 1 FROM sys_role_permission srp , sys_user_role sur WHERE sp.id = srp.permission_id and srp.role_id = sur.role_id and sur.uid = ?)"
	result := getDbFromContext(context).Raw(sql, constants.DbNormalStatus, uid).Scan(&permissions)
	return permissions, result.Error
}

func (s *SysPermissionDao) FindValidPermissions(context context.Context) ([]*SysPermissionPO, error) {
	var permissions []*SysPermissionPO
	result := getDbFromContext(context).Where("status = ?", constants.DbNormalStatus).Find(&permissions)
	return permissions, result.Error
}
