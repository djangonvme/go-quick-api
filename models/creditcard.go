package models

import "github.com/jinzhu/gorm"

type CreditCard struct {
	gorm.Model
	UserID  uint
	Number  string
}
