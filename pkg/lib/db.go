package lib

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"

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

func NewDb(cfg CfgDatabase, dbLogger *DBLogger) (*gorm.DB, error) {
	connArgs := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=%s", cfg.User, cfg.Password, cfg.Host, cfg.Database, "Local")
	db, err := gorm.Open(cfg.Schema, connArgs)
	if err != nil {
		return nil, fmt.Errorf("couldn't connect to database, check your connect args in config.ini, errM: %s", err.Error())
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

	db.Callback().Create().Replace("gorm:update_time_stamp", func(scope *gorm.Scope) {
		scope.SetColumn("CreatedAt", time.Now())
		scope.SetColumn("UpdatedAt", time.Now())
	})
	db.Callback().Update().Replace("gorm:update_time_stamp", func(scope *gorm.Scope) {
		scope.SetColumn("UpdatedAt", time.Now())
	})
	// sql 写入日志 或控制台， 二选一
	if dbLogger != nil {
		db.SetLogger(dbLogger)
	}
	return db, nil
}

//
type DBLogger struct {
	baseLog *logrus.Logger
}

func (d *DBLogger) Print(v ...interface{}) {
	switch v[0] {
	case "sql":
		// sql2 := gorm.LogFormatter(v)
		sql := fmt.Sprintf("%v", v[3])
		// sql = fmt.Sprintf(strings.ReplaceAll(sql, `=?`, `=%v`),v[4])
		d.baseLog.WithFields(
			logrus.Fields{
				"trace":   v[1],
				"latency": fmt.Sprintf("%v", v[2]),
				"sql":     sql,
				//"sql2":     sql2,
				"values": v[4],
				"rows":   v[5],
				"type":   "sql",
			},
		).Infof("sql: %v|%v", v[2], v[5])
	case "log":
		d.baseLog.WithFields(map[string]interface{}{"module": "gorm", "type": "log"}).Print(v[2])
	}
}
