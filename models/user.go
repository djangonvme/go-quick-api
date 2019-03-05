package models

import (
	"github.com/jinzhu/gorm"
	"time"
	"database/sql"
	. "github.com/jangozw/gintest/database"
)


//这里是model 和 表的定义，根据这个结构体，orm执行直接可以创建表 ``里面的是struct tag 用来约束标记什么
type User struct {
	gorm.Model
	Birthday     time.Time
	Age          int
	Name         string  `gorm:"size:255"`       // string默认长度为255, 使用这种tag重设。
	Num          int     `gorm:"AUTO_INCREMENT"` // 自增

	CreditCard        CreditCard      // One-To-One (拥有一个 - CreditCard表的UserID作外键)
	Emails            []Email         // One-To-Many (拥有多个 - Email表的UserID作外键)

	BillingAddress    Address         // One-To-One (属于 - 本表的BillingAddressID作外键)
	BillingAddressID  sql.NullInt64

	ShippingAddress   Address         // One-To-One (属于 - 本表的ShippingAddressID作外键)
	ShippingAddressID int

	IgnoreMe          int `gorm:"-"`   // 忽略这个字段
	Languages         []Language `gorm:"many2many:user_languages;"` // Many-To-Many , 'user_languages'是连接表

}

func (data *User)AddUser() error {
	data.Name = "花粥"
	//data.Age = 25
	data.ShippingAddressID = 100
	data.CreatedAt = time.Now()
	data.Birthday = time.Now()
	r := Db.Create(data)
	return r.Error
}

func (User) List(page int, size int) (lists []User) {
	o := (page - 1) * size
	r := Db.Offset(o).Limit(size).Find(&lists)
	if r.Error != nil {
		panic(r.Error)
	}
	return
}

func (User) Detail(id int) (detail User)  {
	r := Db.First(&detail, id)
	if r.Error != nil {
		panic(r.Error)
	}
	return
}
