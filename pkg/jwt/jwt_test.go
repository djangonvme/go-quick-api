package jwt

import (
	"fmt"
	"testing"
	"time"
)

type testPayload struct {
	UserID int
	Name   string
}

func TestGenerateJwtToken(t *testing.T) {
	token, err := GenerateToken(DefaultSecret, "100", 1)
	if err != nil {
		fmt.Println("gen err: ", err.Error())
		return
	}
	fmt.Println("generate: ", token)

	time.Sleep(time.Second * 2)

	jwtClaim, err := ParseToken(DefaultSecret, token)

	if err != nil {
		fmt.Println("parse err: ", err.Error())
		return
	}
	fmt.Println("decode: ", jwtClaim)
	fmt.Println("decode userId: ", jwtClaim.Id)

}
