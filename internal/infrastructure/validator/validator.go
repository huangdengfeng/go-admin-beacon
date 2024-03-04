package validator

import "github.com/go-playground/validator/v10"

var Validator = validator.New(validator.WithRequiredStructEnabled())

func Struct(any any) error {
	return Validator.Struct(any)
}
func Var(filed any, tag string) error {
	return Validator.Var(filed, tag)
}
