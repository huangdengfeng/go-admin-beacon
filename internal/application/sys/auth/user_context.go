package auth

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-admin-beacon/internal/domain/sys"
)

const UserKey = "userKey"

// ContextWithUser 用户上下文
func ContextWithUser(c context.Context, details *sys.UserDetailsVO) context.Context {
	return context.WithValue(c, UserKey, details)
}

// GetUserFromContext 获取用户
func GetUserFromContext(c context.Context) *sys.UserDetailsVO {
	value := c.Value(UserKey)
	if value != nil {
		return value.(*sys.UserDetailsVO)
	} else {
		return nil
	}
}

// setGinUserContext gin 上下文放入用户信息
func setGinUserContext(c *gin.Context, details *sys.UserDetailsVO) {
	c.Set(UserKey, details)
}

// SetUserToContext 从gin 的上下文中提取用户信息，放入上下文
func SetUserToContext(c *gin.Context) context.Context {
	if value, exists := c.Get(UserKey); exists {
		return ContextWithUser(c, value.(*sys.UserDetailsVO))
	}
	return c
}
