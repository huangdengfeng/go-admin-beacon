package dao

import (
	"go-admin-beacon/internal/infrastructure/errors"
	"gorm.io/gorm"
	"time"
)

// 自定义表名
func (s SysUserPO) TableName() string {
	return "sys_user"
}

type SysUserPO struct {
	//Uid      int32 `gorm:"primaryKey"`
	Uid        int32
	UserName   string
	Password   string
	SecretKey  string
	Name       string
	Mobile     string
	Email      string
	Photo      string
	Status     int8
	CreateBy   int32
	CreateTime time.Time
	UpdateBy   int32
	UpdateTime time.Time
	Remark     string
}

type SysUserDao struct {
	*dao
}

var SysUserDaoInstance = &SysUserDao{&dao{getDb}}

func (s *SysUserDao) FindByUid(uid int32) (*SysUserPO, error) {
	var po SysUserPO
	result := s.db().Take(&po, "uid = ?", uid)
	if nil == result.Error {
		return &po, nil
	}
	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	} else {
		return nil, errors.Newf(errors.SqlError, result.Error.Error())
	}
}

func (s *SysUserDao) FindByUserName(userName string) (*SysUserPO, error) {
	var po SysUserPO
	result := s.db().Take(&po, "user_name = ?", userName)
	if nil == result.Error {
		return &po, nil
	}
	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	} else {
		return nil, errors.Newf(errors.SqlError, result.Error.Error())
	}
}
