package response

import (
	"go-admin-beacon/internal/infrastructure/errors"
	"reflect"
)

// Response 通用响应体
type Response struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

var customErrType = reflect.TypeOf(&errors.Error{})

// DefaultSuccess 成功
var DefaultSuccess = &Response{Code: errors.Success.Code, Msg: errors.Success.Msg}

// Success 成功
func Success(data any) *Response {
	return &Response{Code: errors.SuccessCode, Msg: errors.SuccessMsg, Data: data}
}

// ErrorWithCodeMsg 失败
func ErrorWithCodeMsg(code int32, msg string) *Response {
	return &Response{Code: code, Msg: msg}
}

// Error 失败
func Error(err error) *Response {
	// 是自定义错误类型
	if e, ok := err.(errors.Error); ok {
		return &Response{Code: e.Code, Msg: e.Msg}
	}
	formattedErr := errors.Newf(errors.Unknown, err.Error())
	return &Response{Code: formattedErr.Code, Msg: formattedErr.Msg}
}
