package errors

var Success = New(0, "success")
var Unknown = New(1000, "system errors:%s")
var BadArgs = New(1001, "bad args:%s")
var SqlError = New(1002, "sql errors:%s")

var UserPasswdWrong = New(2005, "用户名或密码错误")