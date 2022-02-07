package singleton

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 这个不能删
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gitlab.com/task-dispatcher/config"
	"gitlab.com/task-dispatcher/model"
	"time"
)

type DB struct {
	*gorm.DB
}

func NewDB(cfg *config.Config, logger *Logger) (*DB, error) {
	args := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=%s", cfg.MySQL.User, cfg.MySQL.Password, cfg.MySQL.Host, cfg.MySQL.DbName, "Local")
	db, err := gorm.Open("mysql", args)
	if err != nil {
		return nil, errors.Errorf("couldn't connect to mysql, check your connect args in config.toml, errM: %s", err.Error())
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
	db.SetLogger(logger.NewDbLogger())

	db.Set("gorm:table_options", "ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci").AutoMigrate(&model.LotusCommit2Task{})
	db.Set("gorm:table_options", "ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci").AutoMigrate(&model.LotusCommit2TaskWorker{})

	return &DB{
		db,
	}, nil
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
