package model

import (
	"fmt"
	"testing"

	"github.com/jangozw/go-quick-api/pkg/app"
	"github.com/stretchr/testify/assert"
)

func TestDb(t *testing.T) {
	app.Init()
	user := &User{}
	err := app.Db.Where("id=?", 1).First(&user).Error
	fmt.Println("user:", user.ID)
	assert.Nil(t, err)
}
