package dao

import (
	"go-admin-beacon/internal/infrastructure/constants"
	"go-admin-beacon/internal/infrastructure/errors"
	"gorm.io/gorm"
	"time"
)

type SysUserPO struct {
	Uid int32 `gorm:"primaryKey"`
	// Not null
	UserName string
	Password string
	// Not null
	SecretKey string
	// Not null
	Name   string
	Mobile string
	Email  string
	Photo  string
	// Not null
	Status     int8
	CreateUser int32
	CreateTime time.Time
	UpdateUser int32
	UpdateTime time.Time
	Remark     string
}

var sysUserAllowedFields = []string{"createTime"}

// SysUserPOCondition 多条件查询
type SysUserPOCondition struct {
	UserName        string
	FuzzyName       string
	Name            string
	Status          *int8
	IncludeSysAdmin bool
	OrderBy         string
}

// TableName 自定义表名
func (s SysUserPO) TableName() string {
	return "sys_user"
}

type SysUserDao struct {
	*dao
}

var SysUserDaoInstance = &SysUserDao{&dao{getDb}}

func (s *SysUserDao) FindByPage(condition *SysUserPOCondition, page int, pageSize int) (*[]SysUserPO, *int64, error) {
	db := s.db().Limit(pageSize).Offset(pageSize * (page - 1))
	if condition.UserName != "" {
		db.Where("user_name = ?", condition.UserName)
	}
	if condition.FuzzyName != "" {
		db.Where("user_name LIKE ?", "%"+condition.UserName+"%")
	}
	if condition.Name != "" {
		db.Where("name = ?", condition.Name)
	}
	if condition.Status != nil {
		db.Where("status = ?", condition.Status)
	}
	if !condition.IncludeSysAdmin {
		db.Not("uid = ?", constants.SuperAdminUserId)
	}

	if condition.OrderBy != "" {
		if convertedOrder, err := checkAndConvertOrder(condition.OrderBy, sysUserAllowedFields); err != nil {
			return nil, nil, err
		} else {
			db.Order(convertedOrder)
		}
	}
	var users []SysUserPO
	if result := db.Find(&users); result.Error != nil {
		return nil, nil, result.Error
	}
	var total int64
	if result := db.Count(&total); result.Error != nil {
		return nil, nil, result.Error
	}
	return &users, &total, nil
}

func (s *SysUserDao) FindByUid(uid int32) (*SysUserPO, error) {
	var po SysUserPO
	// .Clauses(clause.Locking{Strength: "UPDATE"})
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

func (s *SysUserDao) Save(po *SysUserPO) (*int32, error) {
	var uid *int32
	err := s.db().Transaction(func(tx *gorm.DB) error {
		result := tx.Create(po)
		if result.Error != nil {
			return errors.Newf(errors.SqlError, result.Error.Error())
		}
		if result.RowsAffected != 1 {
			return errors.Newf(errors.RowsAffectedNotMatch, result.RowsAffected)
		}
		uid = &po.Uid
		return nil
	})
	return uid, err
}

func (s *SysUserDao) Update(po *SysUserPO) error {
	err := s.db().Transaction(func(tx *gorm.DB) error {
		result := tx.Save(po)
		if result.Error != nil {
			return errors.Newf(errors.SqlError, result.Error.Error())
		}
		if result.RowsAffected != 1 {
			return errors.Newf(errors.RowsAffectedNotMatch, result.RowsAffected)
		}
		return nil
	})
	return err
}
