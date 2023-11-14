package sys

// 角色列表

import (
	"context"
	"go-admin-beacon/internal/domain/dao"
	"go-admin-beacon/internal/infrastructure/response"
)

type RoleListQry struct {
	OnlyValid bool `json:"onlyValid"`
}

type RoleCO struct {
	Id     int32  `json:"id"`
	Code   string `json:"code"`
	Name   string `json:"name"`
	Status int8   `json:"status"`
}

type roleListQryExe struct {
	sysRoleDao *dao.SysRoleDao
}

func NewRoleListQryExe() *roleListQryExe {
	return &roleListQryExe{dao.SysRoleDaoInstance}
}

func (e *roleListQryExe) Execute(context context.Context, qry *RoleListQry) (*response.Response, error) {
	var roles []*dao.SysRolePO
	var err error
	if qry.OnlyValid {
		roles, err = e.sysRoleDao.FindValidRoles(context)
		if err != nil {
			return nil, err
		}
	} else {
		roles, err = e.sysRoleDao.FindAllRoles(context)
		if err != nil {
			return nil, err
		}
	}

	var cos = make([]*RoleCO, 0, len(roles))
	for _, role := range roles {
		cos = append(cos, &RoleCO{
			Id:     role.Id,
			Code:   role.Code,
			Name:   role.Name,
			Status: role.Status,
		})
	}
	return response.Success(cos), nil
}
