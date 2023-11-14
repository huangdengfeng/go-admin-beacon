package constants

//public static final Integer SUPER_ADMIN_USER_ID = 1;
//
//public static final String COMMA = ",";
//public static final String UNDERLINE = "_";

const SuperAdminUserId = 1
const Comma = ","
const Underline = "_"
const DbNormalStatus = 1
const DbInvalidStatus = 2

// RolePrefix 兼容java spring security 写法
const RolePrefix = "ROLE_"

func IsSuperAdmin(uid int32) bool {
	return uid == SuperAdminUserId
}
