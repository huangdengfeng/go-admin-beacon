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
	sql := "SELECT * FROM sys_role sr WHERE sr.status = ? AND EXISTS(SELECT 1 FROM sys_user_role sur WHERE sur.role_id = sr.id AND sur.uid = ?)"
	result := getDbFromContext(context).Raw(sql, constants.DbNormalStatus, uid).Scan(&roles)
	return roles, result.Error
}

func (s *SysRoleDao) FindValidRoles(context context.Context) ([]*SysRolePO, error) {
	var pos []*SysRolePO
	result := getDbFromContext(context).Where("status = ?", constants.DbNormalStatus).Find(&pos)
	return pos, result.Error
}

func (s *SysRoleDao) FindAllRoles(context context.Context) ([]*SysRolePO, error) {
	var pos []*SysRolePO
	result := getDbFromContext(context).Find(&pos)
	return pos, result.Error
}
