package models

import (
	"errors"
	"gin-api-common/databases"
	"gin-api-common/utils"
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
		Password: utils.Sha256(pwd),
	}
	return databases.Db.Create(&user).Error
}
