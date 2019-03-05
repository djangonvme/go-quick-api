package models

import (
	"time"
	"github.com/jangozw/gintest/helper"
	. "github.com/jangozw/gintest/database"
	//"fmt"
)

type Token struct {
	Id int `gorm:"AUTO_INCREMENT"`
	UserId int `gorm:"not null"`
	Token string `gorm:"type:char(32);unique;not null"`
	ExpiredAt time.Time
	UpdatedAt time.Time
	//one-to-one关系, token 表的user_id 关联 user 表的id
	User  User `gorm:"ForeignKey:UserId;AssociationForeignKey:Id"`
}

func (t Token) IsTokenExpired() bool {
	return helper.IsExpired(t.ExpiredAt)
}

func (t Token) Add(uid int, expired int) error{
	et := t.GetByUid(uid)
	if et.Id != 0 {
		return nil
	}
	m := Token{}
	m.UserId = uid
	m.Token = helper.GetMd5RandString()
	m.UpdatedAt = time.Now()
	m.ExpiredAt = helper.GetExpiredAt(expired)
	r:= Db.Create(&m)
	return r.Error
}

func (t Token)GetByUid(uid int) Token {
	et := Token{}
	Db.Where("user_id = ?", uid).First(&et)
	return et
}

func (Token)GetByToken(t string) Token {
	et := Token{}
	Db.Where("token = ?", t).First(&et)
	return et
}

