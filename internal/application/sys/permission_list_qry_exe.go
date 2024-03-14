package sys

import (
	"context"
	"go-admin-beacon/internal/application/sys/auth"
	"go-admin-beacon/internal/domain/dao"
	"go-admin-beacon/internal/infrastructure/errors"
	"go-admin-beacon/internal/infrastructure/response"
)

type PermissionListQry struct {
	OnlyValid bool `json:"onlyValid"`
}

type PermissionCO struct {
	Id       int32           `json:"id"`
	Code     string          `json:"code"`
	Name     string          `json:"name"`
	Status   int8            `json:"status"`
	Children []*PermissionCO `json:"children"`
}
type permissionListQryExe struct {
	sysPermissionDao *dao.SysPermissionDao
}

func NewPermissionListQryExe() *permissionListQryExe {
	return &permissionListQryExe{dao.SysPermissionDaoInstance}
}

func (e permissionListQryExe) Execute(ctx context.Context, qry *PermissionListQry) (*response.Response, error) {
	userDetailsVO := auth.GetUserFromContext(ctx)
	if nil == userDetailsVO {
		return nil, errors.LoginSessionInvalid
	}

	pos, err := e.sysPermissionDao.FindPermissionsByUid(ctx, userDetailsVO.UserId)
	if err == nil {
		return nil, err
	}

	cos := make([]*PermissionCO, 0, len(pos))
	for _, po := range pos {
		// root 节点
		if po.ParentId == nil {
			cos = append(cos, convert(po))
		}
	}
	for _, co := range cos {
		buildTree(co, pos)
	}
	return response.Success(cos), nil
}

func buildTree(parent *PermissionCO, all []*dao.SysPermissionPO) {
	for _, po := range all {
		if po.ParentId != nil && *po.ParentId == parent.Id {
			co := convert(po)
			parent.Children = append(parent.Children, co)
			buildTree(co, all)
		}
	}
}

func convert(po *dao.SysPermissionPO) *PermissionCO {
	return &PermissionCO{
		Id:     po.Id,
		Code:   po.Code,
		Name:   po.Name,
		Status: po.Status,
	}
}
