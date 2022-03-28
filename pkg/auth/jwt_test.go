package auth

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testPayload struct {
	UserID int
	Name   string
}

func TestGenerateJwtToken(t *testing.T) {
	user := testPayload{
		UserID: 100,
		Name:   "hello",
	}
	token, err := GenerateJwtToken("123456", 600, user)
	fmt.Println(err)
	// assert.Nil(t, err)
	fmt.Println(token)
	jwtPayload, err := ParseJwtToken(token, "123456")
	assert.Nil(t, err)
	if err != nil {
		return
	}
	user2 := testPayload{}
	err = jwtPayload.ParseUser(&user2)
	assert.Nil(t, err)
	fmt.Println(user2, err)
}
