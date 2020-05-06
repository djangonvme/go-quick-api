package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jangozw/gin-api-common/libs"
	"time"
)

//jwt 中用户自定义携带参数部分
type AppPayload struct {
	UserId    int64  `json:"uid"`
	UserToken string `json:"uen"`
}

type JwtCustomClaims struct {
	AppPayload
	jwt.StandardClaims
}

//生成一个jwt token
func GenerateJwtToken(p AppPayload, jwtSecret string) (jwtToken string, err error) {
	secret := []byte(jwtSecret)
	c := JwtCustomClaims{
		AppPayload: p,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + libs.Config.GetTokenExpireSeconds(),
			Issuer:    "test", //签发者
		},
	}
	j := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	jwtToken, err = j.SignedString(secret)
	return
}

//解析一个jwt token
func ParseJwtToken(jwtToken string, jwtSecret string) (*JwtCustomClaims, error) {
	secret := []byte(jwtSecret)
	token, err := jwt.ParseWithClaims(jwtToken, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*JwtCustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return claims, err
	}
}
