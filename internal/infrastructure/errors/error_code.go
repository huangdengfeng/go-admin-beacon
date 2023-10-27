package errors

var Success = New(0, "success")
var Unknown = New(1000, "system errors:%s")
var BadArgs = New(1001, "bad args:%s")
var SqlError = New(1002, "sql errors:%s")
var RowsAffectedNotMatch = New(1003, "rows affected:%d")
var OrderByNotAllowed = New(1004, "orderBy not allowed")

var UserPasswdWrong = New(2005, "用户名或密码错误")
