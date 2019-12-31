package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/jangozw/gin-api-common/libs"
	"time"
)

//签发token的签名秘钥
func getJwtSecret() (b []byte, err error) {
	if s, err := libs.Config.Get("encrypt", "jwt_secret"); err != nil {
		return b, errors.New("couldn't get the config key : jwt_secret")
	} else {
		if len(s) == 0 {
			return b, errors.New("jwt_secret len too short")
		}
		return []byte(s), nil
	}
}

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
func GenerateJwtToken(p AppPayload) (jwtToken string, err error) {
	var secret []byte
	if secret, err = getJwtSecret(); err != nil {
		return
	}
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
func ParseJwtToken(jwtToken string) (*JwtCustomClaims, error) {
	secret, err := getJwtSecret()
	if err != nil {
		return nil, err
	}
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
