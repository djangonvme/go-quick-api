package models

import (
	"errors"
	"fmt"
	"github.com/jangozw/gin-api-common/configs"
	"github.com/jangozw/gin-api-common/databases"
	"github.com/jangozw/gin-api-common/utils"
)

// User 用户表
type User struct {
	Model
	Name     string // 姓名
	Mobile   string `gorm:"index"` // 手机号
	Password string // 密码
	Status   int8   // 状态
}

//
func AddUser(name, mobile, pwd string) error {
	var total int
	if err := databases.Db.Model(&User{}).Where("mobile=?", mobile).Count(&total).Error; err != nil {
		return err
	}
	if total > 0 {
		return errors.New("account already exist")
	}
	user := User{
		Name:     name,
		Mobile:   mobile,
		Password: MakeUserPwd(pwd),
	}
	fmt.Println("insert User:", user.Password)
	return databases.Db.Create(&user).Error
}

func FindUserByMobile(mobile string) (user User, err error) {
	if err = databases.Db.Where("mobile=?", mobile).First(&user).Error; err != nil {
		return
	}
	return user, nil
}

func MakeUserPwd(input string) string {
	aesSecret, _ := configs.Get("encrypt", "aes_secret")
	return utils.Sha256(input + aesSecret)
}
func (m *User) CheckPwd(input string) bool {
	aesSecret, _ := configs.Get("encrypt", "aes_secret")
	return m.Password == utils.Sha256(input+aesSecret)
}
