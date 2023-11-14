package auth

import (
	"fmt"
	"testing"
)

func f(i **int) {
	var j int = 10
	*i = &j
}
func Test_checkPermission(t *testing.T) {
	var a *int
	f(&a)
	fmt.Println(*a)
}

func Test_checkRole(t *testing.T) {
}
