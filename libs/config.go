package libs

import (
	"flag"
	"fmt"
	"github.com/Unknwon/goconfig"
	"log"
	"os"
	"strconv"
)

//use config
var Config *config

var configFile string

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

// read config file from console args
func init() {
	cmdArgsConfig := flag.String("config", "config.ini", "config file path")
	flag.Parse()
	configFile = *cmdArgsConfig
	if ok, err := isPathExists(*cmdArgsConfig); err != nil {
		log.Println(fmt.Sprintf("init config err: %s", err.Error()))
		os.Exit(0)
	} else if !ok {
		log.Println(fmt.Sprintf("config file %s is not exists, you can start with arg -config={path}", *cmdArgsConfig))
		os.Exit(0)
	}
	fmt.Println("Init config, configFile is", configFile)
	if c, err := goconfig.LoadConfigFile(configFile); err != nil {
		fmt.Println(fmt.Sprintf("Couldn't load config file %s, %s", configFile, err.Error()))
		os.Exit(0)
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

	if err := mkPathIfNotExists(logDir); err != nil {
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
