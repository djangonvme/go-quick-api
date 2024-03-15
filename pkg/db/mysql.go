package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 这个不能删
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gitlab.com/qubic-pool/model"
	"time"
)

var Instance *gorm.DB

func InitDB(host, user, pwd, dbName string, logger *logrus.Logger) error {
	if Instance != nil {
		return nil
	}
	args := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=%s", user, pwd, host, dbName, "Local")
	db, err := gorm.Open("mysql", args)
	if err != nil {
		return errors.Errorf("couldn't connect to mysql, check your connect args in config.toml, errM: %s", err.Error())
	}

	// config gorm db
	// 全局设置表名不可以为复数形式。
	db.SingularTable(true)
	// prevent no where cause update/delete
	db.BlockGlobalUpdate(true)
	// enable log
	db.LogMode(true)
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(3600 * time.Second)

	// sql 写入日志 或控制台， 二选一
	db.SetLogger(&Logger{logger})
	// 自动创建更新表
	model.AutoMigrate(db)
	Instance = db
	return nil
}

type Logger struct {
	*logrus.Logger
}

func (d *Logger) Print(v ...interface{}) {
	switch v[0] {
	case "sql":
		d.WithFields(logrus.Fields{}).Infof("[sql] %v ", gorm.LogFormatter(v...))
	case "log":
		d.WithFields(map[string]interface{}{"module": "gorm", "type": "log"}).Print(v[2])
	}
}
