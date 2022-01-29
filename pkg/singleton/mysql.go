package singleton

import (
	"fmt"
	"gitlab.com/task-dispatcher/config"
	"time"

	"github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql" // 这个不能删
	"github.com/jinzhu/gorm"
)

type DB struct {
	*gorm.DB
}

func NewDB(cfg *config.Config, logger *Logger) (*DB, error) {
	args := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=%s", cfg.MySQL.User, cfg.MySQL.Password, cfg.MySQL.Host, cfg.MySQL.DbName, "Local")
	db, err := gorm.Open("mysql", args)
	if err != nil {
		return nil, fmt.Errorf("couldn't connect to mysql, check your connect args in config.toml, errM: %s", err.Error())
	}
	dbs := &DB{db}
	// config gorm db
	// 全局设置表名不可以为复数形式。
	dbs.SingularTable(true)
	// prevent no where cause update/delete
	dbs.BlockGlobalUpdate(true)
	// enable log
	dbs.LogMode(true)
	dbs.DB.DB().SetMaxIdleConns(20)
	dbs.DB.DB().SetMaxOpenConns(100)
	dbs.DB.DB().SetConnMaxLifetime(3600 * time.Second)

	dbs.Callback().Create().Replace("gorm:update_time_stamp", func(scope *gorm.Scope) {
		scope.SetColumn("CreatedAt", time.Now())
		scope.SetColumn("UpdatedAt", time.Now())
	})
	dbs.Callback().Update().Replace("gorm:update_time_stamp", func(scope *gorm.Scope) {
		scope.SetColumn("UpdatedAt", time.Now())
	})
	// sql 写入日志 或控制台， 二选一
	dbs.SetLogger(logger.NewDbLogger())
	return dbs, nil
}

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
