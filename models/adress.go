package models

import (
	"database/sql"
	. "github.com/jangozw/gintest/database"
)

type Address struct {
	ID       int
	Address1 string         `gorm:"not null;unique"` // 设置字段为非空并唯一
	Address2 string         `gorm:"type:varchar(100);unique"`
	Post     sql.NullString `gorm:"not null"`
}

func (Address) CreateTable() {
	table := &Address{}// this is the table
	if !Db.HasTable(table) {
		if err := Db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").
			CreateTable(table).Error; err != nil {
			panic(err)
		}
	}
}



