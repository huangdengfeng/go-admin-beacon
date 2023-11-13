package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go-admin-beacon/internal/domain/sys"
	"go-admin-beacon/internal/infrastructure/config"
	"go-admin-beacon/internal/infrastructure/errors"
	"go-admin-beacon/internal/infrastructure/response"
	"go-admin-beacon/internal/infrastructure/security"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// 无需验证登录态
var unAuthUris = [...]string{"/public/", "/sys/login/"}

var userDetailsService = sys.UserDetailsServiceInstance

var userCacheMap = make(map[int32]*userCacheItem)

// UserCache 用户信息缓存，如果用redis 则不需要
type userCacheItem struct {
	data       *sys.UserDetailsVO
	expireTime time.Time
}

func Filter(c *gin.Context) {
	path := c.Request.URL.Path
	var checkAuth = true
	for _, uri := range unAuthUris {
		if strings.HasPrefix(path, uri) {
			checkAuth = false
			break
		}
	}
	// url 需要验证登录
	if !checkAuth {
		c.Next()
		return
	}
	// Authorization: Bearer xxxxx
	authorizationHeader := c.Request.Header.Get("Authorization")
	if authorizationHeader == "" {
		c.Status(http.StatusUnauthorized)
		c.Abort()
		return
	}
	token := strings.TrimPrefix(authorizationHeader, "Bearer ")

	tokenInfo, err := security.ParseToken(token, []byte(config.Global.Login.TokenSignKey))
	if err != nil {
		log.Infof("parse security token error %s ", err.Error())
		c.Status(http.StatusUnauthorized)
		c.Abort()
		return
	}
	subject := tokenInfo.Subject
	uid, err := strconv.Atoi(subject)
	if err != nil {
		log.Errorf("token subject:%s,convert error:%s", subject, err.Error())
		c.Status(http.StatusUnauthorized)
		c.Abort()
		return
	}
	// 获取用户信息并将用户信息放入上下文
	details, err := getUserDetails(int32(uid))
	if err != nil {
		log.Errorf("getUserDetails error:%s", err.Error())
		c.JSON(http.StatusOK, errors.LoadUserError)
		c.Abort()
		return
	}
	// 验证token 中checkSum
	sum256 := sha256.Sum256(append([]byte(subject), []byte(details.SecretKey)...))
	if hex.EncodeToString(sum256[:]) != tokenInfo.CheckSum {
		log.Errorf("checkSum error")
		c.Status(http.StatusUnauthorized)
		c.Abort()
		return
	}
	setGinUserContext(c, details)
	c.Next()
	// 统一错误处理
	for _, e := range c.Errors {
		// 无权限，则转换为403
		if e.Err == errors.NoPermission {
			c.Status(http.StatusForbidden)
		} else {
			c.JSON(http.StatusOK, response.Error(e.Err))
		}
		c.Abort()
		return
	}
}

func getUserDetails(uid int32) (*sys.UserDetailsVO, error) {
	userCache, exists := userCacheMap[uid]
	if exists && userCache.expireTime.After(time.Now()) {
		return userCache.data, nil
	}

	if details, err := userDetailsService.GetUserDetails(uid); err == nil {
		// 放入cache
		userCacheMap[uid] = &userCacheItem{
			data:       details,
			expireTime: time.Now().Add(time.Minute * 1),
		}
		return details, nil
	} else {
		return nil, err
	}

}
