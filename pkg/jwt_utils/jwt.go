package jwt_utils

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/xiaoxin1992/kube-admin/pkg/config"
	"time"
)

var jwtSignKey = []byte(config.GetConfig().Jwt.Token)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(username string) (tk string, err error) {
	claims := Claims{
		username,
		jwt.RegisteredClaims{
			// 120分钟过期
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(config.GetConfig().Jwt.ExpireTime) * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    username,
			Subject:   username,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tk, err = token.SignedString(jwtSignKey)
	return
}

func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSignKey, nil
	})
	if err != nil {
		// 解析token失败
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
