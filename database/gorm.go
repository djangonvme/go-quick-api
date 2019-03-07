package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"gin-api-common/config"
)
var Db  *gorm.DB
//自动执行
func init() {
	/*
	config := map[string]string{
		"schema":"mysql",
		//"host":"192.168.0.238",
		"host":"127.0.0.1",
		"port":"3306",
		//"database":"credit_test",
		"database":"test",
		"user":"root",
		"pwd":"123456",
		"timezone":"Asia%2FShanghai",
	}
	GDb, err = gorm.Open("mysql", "root:123456@tcp(192.168.0.238:3306)/test?charset=utf8&parseTime=True&loc=Asia%2FShanghai")
	*/

	config := config.GetMysql()
	var err error
	//沙雕参数，文档也不写全
	Db, err = gorm.Open(config["schema"], config["user"]+":"+config["pwd"]+"@tcp("+config["host"]+")/"+config["database"]+"?charset=utf8&parseTime=True&loc="+config["timezone"])
	if err != nil {
		log.Fatal(err.Error())
	}
}

func GetDb() *gorm.DB {
	return Db
}

func CloseDb()  {
	Db.Close()
}
