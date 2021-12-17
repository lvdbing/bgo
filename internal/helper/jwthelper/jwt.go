package jwthelper

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/lvdbing/bgo/global"
	"github.com/lvdbing/bgo/pkg/utils"
)

type Claims struct {
	Key string
	jwt.StandardClaims
}

func GenerateToken(key string) (string, error) {
	now := time.Now()
	expireTime := now.Add(global.JWTSetting.Expire)
	claims := Claims{
		Key: utils.EncodeMD5(key),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JWTSetting.Issuer,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(getJWTSecret())
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return getJWTSecret(), nil
	})
	if tokenClaims == nil {
		return nil, err
	}

	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
		return claims, nil
	}
	return nil, err
}

func getJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}
