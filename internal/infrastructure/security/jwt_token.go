package security

import (
	"github.com/golang-jwt/jwt/v5"
	"go-admin-beacon/internal/infrastructure/errors"
	"time"
)

// TokenInfo jwt中关键信息
type TokenInfo struct {
	// 主题字段，业务关键信息
	Subject string
	// 用于安全控制，可空
	TokenId string
	// 验证码，因为jwt 内容没有加密，如果泄露的服务器签名密钥，可以任意伪造，该字段提高伪造难度
	CheckSum string
}

func CreateToken(tokenInfo *TokenInfo, expire time.Duration, secretKey []byte) (string, error) {
	info := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       tokenInfo.TokenId,
		"sub":      tokenInfo.Subject,
		"checkSum": tokenInfo.CheckSum,
		"exp":      time.Now().Add(expire).Unix(),
	})
	if token, err := info.SignedString(secretKey); err != nil {
		return "", err
	} else {
		return token, nil
	}
}

// ParseToken 验证token
func ParseToken(accessToken string, secretKey []byte) (*TokenInfo, error) {
	// 过期或者签名不对都会有错误
	parsedToken, err := jwt.Parse(accessToken, func(token *jwt.Token) (any, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}
	if !parsedToken.Valid {
		return nil, errors.AuthNotPass
	}

	tokeInfo := &TokenInfo{}
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok {
		if tokeInfo.Subject, err = claims.GetSubject(); err != nil {
			return nil, err
		}
		if id, ok := claims["id"].(string); ok {
			tokeInfo.TokenId = id
		}

		if checkSum, ok := claims["checkSum"].(string); ok {
			tokeInfo.CheckSum = checkSum
		}
		return tokeInfo, nil
	}

	return tokeInfo, nil
}
