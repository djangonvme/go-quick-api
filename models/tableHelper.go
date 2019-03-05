package models

import (
	. "github.com/jangozw/gintest/database"
	"fmt"
	"github.com/jangozw/gintest/helper"
)

//model 就是models 包的每一个表的定义struct
func CreateTable(model interface{}) {
	//table := &Address{}// this is the table
	if !Db.HasTable(model) {
		if err := Db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").
			CreateTable(model).Error; err != nil {
			panic(err)
		} else {
			fmt.Println("ok, created success!")
		}
	} else {
		fmt.Println("table has areadly exists")
	}
}

func CreateToken()  {
	CreateTable(Token{})
}

func CreateAllTable()  {
	CreateTable(User{})
	CreateTable(Token{})
	CreateTable(Address{})
	CreateTable(Email{})
	CreateTable(CreditCard{})
	CreateTable(Language{})
	//test
	User{}.AddUser()
	u := User{}
	Db.First(&u)
	if u.ID != 0 {
		Token{}.Add(1, 86400*100)
	}


}



