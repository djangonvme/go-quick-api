package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// DefaultSecret Don't change secret if it already used.
const (
	DefaultSecret string = "vyj36KSkrFQNJ7z6QGmFdu5tdDKVzmy6"
	Issuer               = "qubic-pool-yds"
)

// GenerateToken 生成一个jwt token
func GenerateToken(secret string, userId string, expire int64) (jwtToken string, err error) {
	data := jwt.StandardClaims{
		Id:        userId,
		ExpiresAt: time.Now().Unix() + expire,
		Issuer:    Issuer, // 签发者
		IssuedAt:  time.Now().Unix(),
	}
	j := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	if token, err := j.SignedString([]byte(secret)); err != nil {
		return "", errors.New("generate jwt token failed: " + err.Error())
	} else {
		return token, nil
	}
}

// ParseToken 解析一个jwt token
func ParseToken(secret, jwtToken string) (*jwt.StandardClaims, error) {
	if jwtToken == "" {
		return nil, errors.New("empty jwt token")
	}
	token, err := jwt.ParseWithClaims(jwtToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, errors.New("jwt token parsed:" + err.Error())
	}
	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		return claims, nil
	} else {
		return claims, errors.New("jwt decode failed")
	}
}
