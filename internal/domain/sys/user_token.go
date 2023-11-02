package sys

import (
	"crypto/sha256"
	"encoding/hex"
	"go-admin-beacon/internal/domain/dao"
	"go-admin-beacon/internal/infrastructure/config"
	"go-admin-beacon/internal/infrastructure/errors"
	"go-admin-beacon/internal/infrastructure/security"
	"strconv"
	"time"
)

type UserTokenService struct {
	sysUserDao *dao.SysUserDao
}

var UserTokenServiceInstance = &UserTokenService{
	sysUserDao: dao.SysUserDaoInstance,
}

func (s *UserTokenService) Create(userName string) (string, error) {
	po, err := s.sysUserDao.FindByUserName(userName)
	if nil != err {
		return "", err
	}
	// 用户不存在
	if po == nil {
		return "", errors.UserNotExists
	}
	subject := strconv.Itoa(int(po.Uid))
	hash := sha256.Sum256(append([]byte(subject), []byte(po.SecretKey)...))

	tokenInfo := &security.TokenInfo{Subject: subject, CheckSum: hex.EncodeToString(hash[:])}
	var loginConfig = config.Global.Login
	return security.CreateToken(tokenInfo, loginConfig.AccessTokenExpire*time.Minute, []byte(loginConfig.TokenSignKey))
}
