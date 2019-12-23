package databases

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" //这个不能删
	"github.com/jangozw/gin-api-common/configs"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

var Db *gorm.DB

type connectConf struct {
	driver string
	args   string
}

func init() {
	if Db != nil {
		return
	}
	//初始化/conf.ini 中database区指定的数据库信息
	Db = initDatabase("database")
}

func DB() *gorm.DB {
	return Db
}

//load database config info from /conf.ini
func initDatabase(confSection string) *gorm.DB {
	c := getConnectConf(confSection)
	db, err := gorm.Open(c.driver, c.args)
	if err != nil {
		log.Fatal("couldn't connect to database:", c.driver, ":", c.args, err.Error())
	}
	//config gorm db
	db.SingularTable(true) // 全局设置表名不可以为复数形式。
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(3600 * time.Second)
	db.LogMode(true)
	db.Callback().Create().Replace("gorm:update_time_stamp", func(scope *gorm.Scope) {
		scope.SetColumn("CreatedAt", time.Now().Unix())
		scope.SetColumn("UpdatedAt", time.Now().Unix())
	})
	db.Callback().Update().Replace("gorm:update_time_stamp", func(scope *gorm.Scope) {
		scope.SetColumn("UpdatedAt", time.Now().Unix())
	})
	return db
}

//数据库连接配置
func getConnectConf(section string) *connectConf {
	c, err := configs.GetSection(section)
	if err != nil {
		log.Fatal("couldn't get config info by section:" + section)
	}
	return &connectConf{
		driver: c["schema"],
		args:   fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=%s", c["user"], c["pwd"], c["host"], c["port"], c["dbname"], c["timezone"]),
	}
}
