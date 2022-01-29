package model

import (
	"errors"

	"gitlab.com/task-dispatcher/pkg/app"
	"gitlab.com/task-dispatcher/pkg/util"
	"gitlab.com/task-dispatcher/types"
)

const (
	UserStatusNormal    = 1
	UserStatusForbidden = 2
)

// User 用户表
type User struct {
	Model
	Name     string // 姓名
	Mobile   string `gorm:"index"` // 手机号
	Password string // 密码
	Status   int8   // 状态
}

func (m *User) CheckPwd(input string) bool {
	return m.Password == util.Sha256(input)
}

// UserToken 用户token表
type UserToken struct {
	Model
	UserID    int64  `gorm:"index"` // 用户id
	Token     string // token
	ExpiredAt int64  // 过期时间
}

//
func AddUser(name, mobile, pwd string) (User, error) {
	var total int
	if err := app.Db().Model(&User{}).Where("mobile=?", mobile).Count(&total).Error; err != nil {
		return User{}, err
	}
	if total > 0 {
		return User{}, errors.New("account already exist")
	}
	user := User{
		Name:     name,
		Mobile:   mobile,
		Password: makeUserPwd(pwd),
	}
	return user, app.Db().Create(&user).Error
}

func FindUserByMobile(mobile string) (user User, err error) {
	if err = app.Db().Where("mobile=?", mobile).First(&user).Error; err != nil {
		return
	}
	return user, nil
}

func FindUserByID(ID int64) (user User, err error) {
	if err = app.Db().Where("id=?", ID).First(&user).Error; err != nil {
		return
	}
	return user, nil
}

func makeUserPwd(input string) string {
	return util.Sha256(input)
}

func UserList(search types.UserListRequest, pager app.Pager) ([]User, error) {
	var users []User
	query := app.Db().Model(&User{})
	if search.Mobile != "" {
		query = query.Where("mobile=?", search.Mobile)
	}
	err := query.Offset(pager.Offset()).Limit(pager.Limit()).Find(&users).Error
	return users, err
}
