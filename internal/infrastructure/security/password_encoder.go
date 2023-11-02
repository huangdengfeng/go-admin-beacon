package security

import (
	"golang.org/x/crypto/bcrypt"
)

// 轮次强度
const strength = 10

func EncodePassword(password string) (string, error) {
	generated, err := bcrypt.GenerateFromPassword([]byte(password), strength)
	return string(generated), err
}

func PasswordMatches(rawPassword string, encodedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encodedPassword), []byte(rawPassword))
	return err == nil
}
