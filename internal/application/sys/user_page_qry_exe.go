package sys

// 用户分页信息查询

import (
	"context"
	"go-admin-beacon/internal/application/sys/auth"
	"go-admin-beacon/internal/domain/dao"
	"go-admin-beacon/internal/infrastructure/request"
	"go-admin-beacon/internal/infrastructure/response"
	"time"
)

type UserPageQry struct {
	request.PageQry
	Uid       int32  `json:"uid"`
	UserName  string `json:"userName"`
	FuzzyName string `json:"fuzzyName"`
	Status    *int8  `json:"status"`
}

type UserCO struct {
	Uid        int32  `json:"uid"`
	UserName   string `json:"userName"`
	Name       string `json:"name"`
	Mobile     string `json:"mobile"`
	Email      string `json:"email"`
	Photo      string `json:"photo"`
	Status     int8   `json:"status"`
	CreateUser int32  `json:"createUser"`
	CreateTime string `json:"createTime"`
	UpdateUser int32  `json:"updateUser"`
	UpdateTime string `json:"updateTime"`
	Remark     string `json:"remark"`
}

type userPageQryExe struct {
	sysUserDao *dao.SysUserDao
}

func NewUserPageQryExe() *userPageQryExe {
	return &userPageQryExe{dao.SysUserDaoInstance}
}

func (e *userPageQryExe) Execute(context context.Context, qry *UserPageQry) (*response.Response, error) {
	if err := auth.CheckPermission(context, "sys:user:qry"); err != nil {
		return nil, err
	}
	condition := &dao.SysUserPOCondition{
		UserName:  qry.UserName,
		FuzzyName: qry.FuzzyName,
		Status:    qry.Status,
		OrderBy:   qry.OrderBy,
	}

	pos, total, err := e.sysUserDao.FindByPage(condition, qry.Page, qry.PageSize)
	if err != nil {
		return nil, err
	}

	cos := make([]*UserCO, 0, len(pos))

	for _, po := range pos {
		cos = append(cos, &UserCO{
			Uid:        po.Uid,
			UserName:   po.UserName,
			Name:       po.Name,
			Mobile:     po.Mobile,
			Email:      po.Email,
			Photo:      po.Photo,
			Status:     po.Status,
			CreateUser: po.CreateUser,
			CreateTime: po.CreateTime.Format(time.DateTime),
			UpdateUser: po.UpdateUser,
			UpdateTime: po.UpdateTime.Format(time.DateTime),
			Remark:     po.Remark,
		})
	}

	page := &response.Page[*UserCO]{
		Total: total,
		Data:  cos,
	}
	return response.Success(page), nil
}
