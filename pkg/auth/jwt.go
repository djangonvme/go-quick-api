package auth

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// jwt token 加载的信息
type JwtPayload struct {
	User interface{} `json:"user"`
	jwt.StandardClaims
}

func (p *JwtPayload) ParseUser(res interface{}) error {
	if p.User == nil || res == nil {
		return errors.New("jwt user payload is empty")
	}
	by, err := json.Marshal(p.User)
	if err != nil {
		return err
	}
	return json.Unmarshal(by, res)
}

// 生成一个jwt token
func GenerateJwtToken(secret string, expire int64, user interface{}) (jwtToken string, err error) {
	data := JwtPayload{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + expire,
			Issuer:    "issuer", // 签发者
		},
	}
	j := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	if token, err := j.SignedString([]byte(secret)); err != nil {
		return "", errors.New("jwt 生成token失败:" + err.Error())
	} else {
		return token, nil
	}
}

// 解析一个jwt token
func ParseJwtToken(jwtToken string, secret string) (*JwtPayload, error) {
	if jwtToken == "" {
		return nil, errors.New("empty jwt token")
	}
	token, err := jwt.ParseWithClaims(jwtToken, &JwtPayload{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, errors.New("jwt 解析失败:" + err.Error())
	}
	if claims, ok := token.Claims.(*JwtPayload); ok && token.Valid {
		return claims, nil
	} else {
		return claims, errors.New("jwt 解析后验证失败")
	}
}
