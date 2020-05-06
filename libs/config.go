package libs

import (
	"errors"
	"fmt"
	"github.com/Unknwon/goconfig"
	"os"
	"strconv"
)

// Config use config
var Config *config

// 配置文件路径可根据自己环境自由修改, 不影响其他环境部署, 部署时是按照命令行参数传递进来的配置文件路径
// 创建项目的配置文件的软连接 sudo ln -s $(pwd)/config.ini /etc/ginapicommon_config.ini
var ConfigFile = "/etc/ginapicommon_config.ini"

type config struct {
	*goconfig.ConfigFile
}

// required config items
var confRequiredFields = map[string][]string{
	"log":      {"log_dir"},
	"server":   {"listen"},
	"database": {"user", "dbName"},
	"redis":    {"host"},
}

func init() {

	if ok, err := isPathExists(ConfigFile); err != nil {
		panic(fmt.Sprintf("init config err: %s", err.Error()))
	} else if !ok {
		panic(fmt.Sprintf("config file %s is not exists, you can start with arg -config={path} ", ConfigFile))
	}
	fmt.Println("init config, configFile is", ConfigFile)
	if c, err := goconfig.LoadConfigFile(ConfigFile); err != nil {
		panic(fmt.Sprintf("Couldn't load config file %s, %s", ConfigFile, err.Error()))
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
				msg = append(msg, fmt.Sprintf("Error: check config fail, %s in [%s] is required in config file "+ConfigFile, field, section))
			}
		}
	}
	logPath := Config.GetLogDir()
	if ok, err := isPathExists(logPath); err != nil {
		msg = append(msg, fmt.Sprintf("log path err: %s", err.Error()))
	} else if !ok {
		msg = append(msg, fmt.Sprintf("log path not exists: %s", logPath))
	}
	if len(msg) > 0 {
		for _, v := range msg {
			fmt.Println(v)
		}
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
func isPathExists(dirPath string) (bool, error) {
	_, err := os.Stat(dirPath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func mkPathIfNotExists(dirPath string) error {
	exist, err := isPathExists(dirPath)
	if err != nil {
		return err
	}
	if !exist {
		// 创建文件夹 =
		return os.Mkdir(dirPath, 0766)
	}
	return nil
}

//签发token的签名秘钥
func GetJwtSecret() (s string, err error) {
	if s, err := Config.Get("encrypt", "jwt_secret"); err != nil {
		return s, errors.New("couldn't get the config key : jwt_secret")
	} else {
		if len(s) == 0 {
			return s, errors.New("jwt_secret len too short")
		}
		return s, nil
	}
}
