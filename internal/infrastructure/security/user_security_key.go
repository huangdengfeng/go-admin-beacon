package security

import (
	"math/rand"
	"time"
)

const randomCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateUserKey() string {
	var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	length := 32
	result := make([]byte, length)
	for i := range result {
		result[i] = randomCharset[seededRand.Intn(len(randomCharset))]
	}
	return string(result)
}
