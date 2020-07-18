package app

import (
	"fmt"

	"github.com/jangozw/go-quick-api/config"
	"github.com/jangozw/go-quick-api/param"
	"github.com/jangozw/go-quick-api/pkg/lib"
	"github.com/jangozw/go-quick-api/pkg/util"
	"github.com/jinzhu/gorm"
)

// 服务列表
const (
	LogService   = `Logger`
	RedisService = `Redis`
	DbService    = `Db`
)

type loadServiceMap map[string]func() error

// 配置文件路径，启动参数指定
var ConfigFile string

// 编译的app版本等信息
var BuildInfo string

// 解析的配置文件
var Cfg *config.Config

// 数据库
var Db *gorm.DB

// redis
var Redis *lib.Redis

// 日志
var Logger *lib.Logger

// 已经加载的服务
var Loaded []string

// app 需要自动加载的服务配置，用不到的可以注释掉
var serviceMap = loadServiceMap{
	LogService:   LoadLogger,
	RedisService: LoadRedis,
	DbService:    LoadDb,
}

func (m loadServiceMap) keys() []string {
	keys := make([]string, 0)
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// app 包初始化，加载服务失败要panic
// cfg，logger 是系统要求强制加载的，其他服务可选
func Init(services ...string) {
	LoadCfg()
	serviceMap[LogService] = LoadLogger
	if len(services) == 0 {
		services = serviceMap.keys()
	}
	Loaded = make([]string, 0)
	for key, load := range serviceMap {
		if util.InStringSlice(key, services) {
			if err := load(); err != nil {
				panic("app load service fail: " + err.Error())
			}
			Loaded = append(Loaded, key)
		}
	}
	if Logger != nil {
		Logger.Infof("%s init services successfully! loaded services: %s", param.AppName, Loaded)
	}
}

// loadCfg 从启动参数或者项目目录中查找并加载配置文件
func LoadCfg() {
	if Cfg != nil {
		return
	}
	// ConfigFile 在启动时候赋值
	if ConfigFile == "" {
		// 配置文件的启动参数名称 eg: -config=xx.ini
		configFlag := param.ArgConfig
		// 配置文件相对于运行目录的路径
		filename := param.ArgConfigFilename

		err := fmt.Errorf("try to find config file %s failed: %s", filename, "no such file")
		f, _ := util.FindConfigFile(filename, configFlag)
		if f == "" {
			panic(err)
		}
		ok, _ := util.IsPathExists(f)
		if !ok {
			panic(err)
		}
		ConfigFile = f
	}
	Cfg = &config.Config{}
	if err := util.ParseIni(ConfigFile, Cfg); err != nil {
		panic(fmt.Sprintf("Parse config file %s to config.Config object failed: %s", ConfigFile, err.Error()))
	}
}

// 加载日志
func LoadLogger() (err error) {
	if Logger != nil {
		return nil
	}
	LoadCfg()
	module := param.AppName
	Logger, err = lib.NewLogger(Cfg.General.LogDir, module)
	if err == nil {
		fmt.Println("app loaded service Logger successfully!  module=" + module + ",log_dir=" + Cfg.General.LogDir)
	}
	return
}

// 加载Db
func LoadDb() (err error) {
	if Db != nil {
		return nil
	}
	LoadCfg()
	cfgDb := lib.CfgDatabase{
		Schema:   Cfg.Database.Schema,
		Host:     Cfg.Database.Host,
		User:     Cfg.Database.User,
		Password: Cfg.Database.Password,
		Database: Cfg.Database.Database,
	}
	Db, err = lib.NewDb(cfgDb)
	if err == nil {
		fmt.Println("app loaded service Db successfully! ")
	}
	return
}

// 加载redis
func LoadRedis() (err error) {
	if Redis != nil {
		return nil
	}
	LoadCfg()
	cfgRedis := lib.CfgRedis{
		Host:     Cfg.Redis.Host,
		Password: Cfg.Redis.Password,
		DbNum:    Cfg.Redis.DbNum,
	}
	Redis, err = lib.NewRedis(cfgRedis)
	if err == nil {
		fmt.Println("app loaded service Redis successfully! ")
	}
	return
}
