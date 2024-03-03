package errors

var Success = New(0, "success")
var Unknown = New(1000, "system errors:%s")
var BadArgs = New(1001, "bad args:%s")
var SqlError = New(1002, "sql errors:%s")
var RowsAffectedNotMatch = New(1003, "rows affected:%d")
var OrderByNotAllowed = New(1004, "orderBy not allowed")
var NoPermission = New(1005, "没有权限")

var UserNotExists = New(2004, "用户名不存在")
var UserPasswdWrong = New(2005, "用户名或密码错误")
var AuthNotPass = New(2006, "认证不通过")
var LoadUserError = New(2007, "加载用户信息出错，请重试")
var UserNameExists = New(2008, "用户名已存在")
