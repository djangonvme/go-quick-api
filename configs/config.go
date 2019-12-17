package configs

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"strconv"
)

var conf *goconfig.ConfigFile

func init() {
	var err error
	filename := "app.ini"
	conf, err = goconfig.LoadConfigFile(filename)
	if err != nil {
		panic("load config file " + filename + " failed!")
	}
}

//
func GetSection(section string) (map[string]string, error) {
	return conf.GetSection(section)
}

//
func Get(section string, key string) (string, error) {
	return conf.GetValue(section, key)
}

func GetInt(section string, key string) int {
	i, _ := Get(section, key)
	n, _ := strconv.Atoi(i)
	return n
}

//过期时长
func GetTokenExpireSeconds() int64 {
	sec := GetInt("encrypt", "token_expire_seconds")
	fmt.Println("asdfasdf=", sec)
	if sec > 86400*30 || sec <= 0 {
		sec = 600
	}
	return int64(sec)
}
