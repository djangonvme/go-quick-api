package configs

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"os"
	"strconv"
)

const confFile = "app.ini"

var confRequiredFields = map[string][]string{
	"env":      {"log_dir"},
	"server":   {"listen"},
	"database": {"user", "dbname"},
	"redis":    {"redis_host"},
	//"secret":   {"aes_secret"},
}

var conf *goconfig.ConfigFile

func init() {
	if c, err := goconfig.LoadConfigFile(confFile); err != nil {
		panic("Load config file " + confFile + " failed!")
	} else {
		conf = c
	}
	checkRequired()
}
func checkRequired() {
	var msg []string
	for section, val := range confRequiredFields {
		for _, field := range val {
			if _, err := conf.GetValue(section, field); err != nil {
				msg = append(msg, fmt.Sprintf("Error: check config fail, %s in [%s] is required in file %s ", field, section, confFile))
			}
		}
	}
	if len(msg) > 0 {
		for _, v := range msg {
			fmt.Println(v)
		}
		fmt.Println("Exit!")
		os.Exit(0)
	}

	logDir, _ := Get("env", "log_dir")
	if err := MkDirIfNotExists(logDir); err != nil {
		fmt.Println("Error: mkdir " + logDir + " failed, " + err.Error())
		os.Exit(0)
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

//
func GetInt(section string, key string) (int, error) {
	if v, err := Get(section, key); err != nil {
		return 0, err
	} else {
		return strconv.Atoi(v)
	}
}

//过期时长
func GetTokenExpireSeconds() int64 {
	sec, _ := GetInt("encrypt", "token_expire_seconds")
	if sec > 86400*30 || sec <= 0 {
		sec = 600
	}
	return int64(sec)
}

func GetHttpPort() (port int, err error) {
	return GetInt("server", "listen")
}

func GetLogDir() string {
	d, _ := Get("env", "log_dir")
	return d
}

// 判断文件夹是否存在
func IsDirExists(dirPath string) (bool, error) {
	_, err := os.Stat(dirPath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func MkDirIfNotExists(dirPath string) error {
	exist, err := IsDirExists(dirPath)
	if err != nil {
		return err
	}
	if !exist {
		// 创建文件夹 =
		return os.Mkdir(dirPath, 0766)
	}
	return nil
}
