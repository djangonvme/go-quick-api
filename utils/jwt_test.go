package utils

import (
	"fmt"
	"github.com/jangozw/gin-api-common/libs"
	"testing"
)

func init() {
	//flag.String("config", "", "config file path, default: ")

}

func TestGenerateJwtToken(t *testing.T) {
	p := AppPayload{UserId: 122}
	secret, err := libs.GetJwtSecret()
	if err != nil {
		t.Error(err)
		return
	}
	jwtToken, err := GenerateJwtToken(p, secret)
	if err != nil {
		t.Error("jwtToken generate failed!", err.Error())
	}
	fmt.Println(jwtToken)
}

func TestParseJwtToken(t *testing.T) {
	secret, err := libs.GetJwtSecret()
	if err != nil {
		t.Error(err)
		return
	}
	jwtToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEyMiwidWVuIjoiIiwiZXhwIjoxNTg4NzUxNjY4LCJpc3MiOiJ0ZXN0In0.ameQ3G2f-8X19sDHZL7wZOIy5EMS-NINKx33bBP5E4A"
	c, err := ParseJwtToken(jwtToken, secret)
	if err != nil {
		fmt.Println("err:", err.Error())
	} else {
		fmt.Println(c.UserId, c.ExpiresAt, c.Issuer)
	}
}

func TestRandNum(t *testing.T) {
	//fmt.Println(string(RandNum(12)))
	fmt.Println(RandToken(), len(RandToken()))
}

func TestSha256(t *testing.T) {
	fmt.Println(Sha256("123456" + "intimidating$"))
}
