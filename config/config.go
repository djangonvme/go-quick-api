package config

import (
	"github.com/Unknwon/goconfig"
	"fmt"
)

var conf *goconfig.ConfigFile

func init() {
	var err error
	conf, err = goconfig.LoadConfigFile("conf")
	if err != nil {
		panic("read conf.ini错误")
	}
}

func ShowConf()  {
	fmt.Println(conf)
}

func GetMysql() map[string]string {
	sec,_ := conf.GetSection("mysql")
	return sec
}

func GetValue(section string, key string) string {
	v, _ := conf.GetValue(section, key)
	return v
}
