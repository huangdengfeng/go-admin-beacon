package auth

import (
	"context"
	"golang.org/x/exp/slices"
)

// 检查权限

func CheckPermission(context context.Context, code string) bool {
	userDetailsVO := GetUserFromContext(context)
	return nil != userDetailsVO && slices.Contains(userDetailsVO.PermissionCodes, code)
}

func CheckRole(context context.Context, code string) bool {
	userDetailsVO := GetUserFromContext(context)
	return nil != userDetailsVO && slices.Contains(userDetailsVO.RoleCodes, code)
}
