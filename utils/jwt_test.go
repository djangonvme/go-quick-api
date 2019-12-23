package utils

import (
	"fmt"
	"testing"
)

func TestGenerateJwtToken(t *testing.T) {
	p := AppPayload{UserId: 122}
	jwtToken, err := GenerateJwtToken(p)
	if err != nil {
		t.Error("jwtToken generate failed!", err.Error())
	}
	fmt.Println(jwtToken)

}

func TestParseJwtToken(t *testing.T) {
	jwtToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEyMiwiZXhwIjoxNTc2NDYzMzExLCJpc3MiOiJ0ZXN0In0.-nMJZKAWJXOfnXgStJ5bj6Qvll2tv3GY1Ea96avogUE"
	c, err := ParseJwtToken(jwtToken)
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
