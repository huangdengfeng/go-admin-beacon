package auth

import (
	"context"
	"go-admin-beacon/internal/infrastructure/constants"
	"go-admin-beacon/internal/infrastructure/errors"
	"golang.org/x/exp/slices"
)

// 检查权限

func CheckPermission(context context.Context, code string) error {
	userDetailsVO := GetUserFromContext(context)
	if userDetailsVO == nil {
		return errors.LoginSessionInvalid
	}
	if !slices.Contains(userDetailsVO.PermissionCodes, code) {
		return errors.NoPermission
	}
	return nil
}

func CheckRole(context context.Context, code string) (bool, error) {
	userDetailsVO := GetUserFromContext(context)
	if userDetailsVO == nil {
		return false, errors.LoginSessionInvalid
	}
	return nil != userDetailsVO && slices.Contains(userDetailsVO.RoleCodes, constants.RolePrefix+code), nil
}
