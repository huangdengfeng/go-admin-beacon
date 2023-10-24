package errors

import "fmt"

const (
	SuccessCode = 0
	SuccessMsg  = "success"
)

type Error struct {
	Code int32
	Msg  string
}

func New(code int32, msg string) Error {
	return Error{code, msg}
}

func Newf(error Error, args ...any) Error {
	formatted := fmt.Sprintf(error.Msg, args)
	return New(error.Code, formatted)
}
func (err Error) Error() string {
	if err.Code == SuccessCode {
		return SuccessMsg
	}
	return fmt.Sprintf("code %d, msg: %s", err.Code, err.Msg)
}
