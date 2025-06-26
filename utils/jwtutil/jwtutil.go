package jwtutil

import (
	"errors"
	"k8s-platform/config"
	"k8s-platform/utils/logs"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSignKey = []byte(config.JwtSignKey)

// 自定义声明类型
type MyCustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// 生成JWT
func GetToken(username string) (string, error) {
	claims := MyCustomClaims{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(config.JwtExpireTime))), //过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "admin",
			Subject:   "k8s-platform",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(jwtSignKey)
	return ss, err
}

// 解析token
func ParseToken(ss string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(ss, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSignKey, nil
	})

	if err != nil {
		logs.Error(nil, "解析Token失败")
		return nil, err
	}

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, nil // token合法
	} else {
		logs.Warning(nil, "Token不合法")
		return nil, errors.New("token不合法: invalid token")
	}
}
