package lib

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql" // 这个不能删
	"github.com/jinzhu/gorm"
)

type CfgDatabase struct {
	Schema   string
	Host     string
	User     string
	Password string
	Database string
}

func NewDb(cfg CfgDatabase) (*gorm.DB, error) {
	return cfg.NewDb()
}

func (d *CfgDatabase) NewDb() (*gorm.DB, error) {
	connArgs := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=%s", d.User, d.Password, d.Host, d.Database, "Local")
	db, err := gorm.Open(d.Schema, connArgs)
	if err != nil {
		return nil, fmt.Errorf("couldn't connect to database, check your connect args in config.ini, errMsg: %s", err.Error())
	}
	// config gorm db
	db.SingularTable(true) // 全局设置表名不可以为复数形式。
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(3600 * time.Second)
	// log all sql in console
	db.LogMode(true)

	db.Callback().Create().Replace("gorm:update_time_stamp", func(scope *gorm.Scope) {
		scope.SetColumn("CreatedAt", time.Now())
		scope.SetColumn("UpdatedAt", time.Now())
	})
	db.Callback().Update().Replace("gorm:update_time_stamp", func(scope *gorm.Scope) {
		scope.SetColumn("UpdatedAt", time.Now())
	})
	return db, nil
}
