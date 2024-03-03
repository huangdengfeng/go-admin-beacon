package constants

import (
	"fmt"
	"go-admin-beacon/internal/infrastructure/config"
	"sync"
)

const SuperAdminUserId = 1
const Comma = ","
const Underline = "_"
const DbNormalStatus = 1
const DbInvalidStatus = 2

// RolePrefix 兼容java spring security 写法
const RolePrefix = "ROLE_"
const DictSysUserStatus = "sys-user-status"

func IsSuperAdmin(uid int32) bool {
	return uid == SuperAdminUserId
}

// 自定常量处理
var dictMapMapping = make(map[string]string)
var dictInitOnce = sync.Once{}

func GetDictName(dictType string, value any) string {
	keyFmt := "%s" + Underline + "%v"
	dictInitOnce.Do(func() {
		for key, array := range config.AppDictInstance {
			for _, d := range array {
				dictMapMapping[fmt.Sprintf(keyFmt, key, d.Value)] = d.Name
			}
		}
	})
	key := fmt.Sprintf(keyFmt, dictType, value)
	return dictMapMapping[key]
}
