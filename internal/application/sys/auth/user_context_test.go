package auth

import (
	"testing"
)

func TestGetUserFromContext(t *testing.T) {
	var userCacheMap = make(map[int32]*userCacheItem)
	userCacheMap[1] = nil
	t.Logf("userCacheMap:%d", len(userCacheMap))
	var s = make([]string, 0)
	t.Logf("string:%t", s == nil)

}
