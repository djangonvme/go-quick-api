package libs

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"os"
	"strconv"
)

//use config
var Config *config

// it's soft link to this project's root app.ini

const configFile = "/data/GinApiCommon_config.ini"

type config struct {
	*goconfig.ConfigFile
}

// config file path set
//var configFile string
var confRequiredFields = map[string][]string{
	"log":      {"log_dir"},
	"server":   {"listen"},
	"database": {"user", "dbname"},
	"redis":    {"redis_host"},
	//"secret":   {"aes_secret"},
}

func init() {
	if c, err := goconfig.LoadConfigFile(configFile); err != nil {
		panic("Couldn't load config file " + configFile)
	} else {
		Config = &config{c}
	}
	checkRequired()
}

// check required fields in config file
func checkRequired() {
	var msg []string
	for section, val := range confRequiredFields {
		for _, field := range val {
			if _, err := Config.GetValue(section, field); err != nil {
				msg = append(msg, fmt.Sprintf("Error: check config fail, %s in [%s] is required in config file "+configFile, field, section))
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

	logDir := Config.GetLogDir()

	if err := mkDirIfNotExists(logDir); err != nil {
		fmt.Println("Error: mkdir " + logDir + " failed, " + err.Error())
		os.Exit(0)
	}
}

//
func (c *config) Section(section string) (map[string]string, error) {
	return c.GetSection(section)
}

//
func (c *config) Get(section string, key string) (string, error) {
	return c.GetValue(section, key)
}

//
func (c *config) GetInt(section string, key string) (int, error) {
	if v, err := c.Get(section, key); err != nil {
		return 0, err
	} else {
		return strconv.Atoi(v)
	}
}

//过期时长
func (c *config) GetTokenExpireSeconds() int64 {
	sec, _ := c.GetInt("encrypt", "token_expire_seconds")
	if sec > 86400*30 || sec <= 0 {
		sec = 600
	}
	return int64(sec)
}

func (c *config) GetHttpPort() (port int, err error) {
	return c.GetInt("server", "listen")
}

func (c *config) GetLogDir() string {
	d, _ := c.Get("log", "log_dir")
	return d
}

// 判断文件夹是否存在
func isDirExists(dirPath string) (bool, error) {
	_, err := os.Stat(dirPath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func mkDirIfNotExists(dirPath string) error {
	exist, err := isDirExists(dirPath)
	if err != nil {
		return err
	}
	if !exist {
		// 创建文件夹 =
		return os.Mkdir(dirPath, 0766)
	}
	return nil
}
