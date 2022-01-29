package app

import (
	"gitlab.com/task-dispatcher/config"
	"gitlab.com/task-dispatcher/pkg/singleton"
)

var BuildInfo string
var CfgInstance *config.Config
var DBInstance *singleton.DB
var RedisInstance *singleton.RedisClient
var LoggerInstance *singleton.Logger
var dLocker LockerIf

func Redis() *singleton.RedisClient {
	if RedisInstance == nil {
		panic("RedisInstance is nil")
	}
	return RedisInstance
}

func Db() *singleton.DB {
	if DBInstance == nil {
		panic("DBInstance is nil")
	}
	return DBInstance
}
func Log() *singleton.Logger {
	if LoggerInstance == nil {
		panic("LoggerInstance is nil")
	}
	return LoggerInstance
}

func Cfg() *config.Config {
	if CfgInstance == nil {
		panic("CfgInstance is nil")
	}
	return CfgInstance
}
func Locker() LockerIf {
	if dLocker != nil {
		return dLocker
	}
	if RedisInstance != nil {
		return RedisInstance
	} else {
		dLocker = NewProcessLocker()
		return dLocker
	}
}

/*func LoadLogger() (err error) {
	module := types.AppName
	Logger, err = singleton.NewLogger(config.Cfg.General.LogDir, module)
	if err == nil {
		Logger.Infof("app loaded service Logger successfully!  module=" + module + ",log_dir=" + config.Cfg.General.LogDir)
	}
	return
}*/

/*func LoadDb() (err error) {
	if Db != nil {
		return nil
	}
	cfgDb := singleton.CfgDatabase{
		Schema:   config.Cfg.Database.Schema,
		Host:     config.Cfg.Database.Host,
		User:     config.Cfg.Database.User,
		Password: config.Cfg.Database.Password,
		Database: config.Cfg.Database.Database,
	}
	var dbLogger *singleton.DBLogger

	if Logger != nil {
		dbLogger = Logger.NewDbLogger()
	}

	Db, err = singleton.NewDb(cfgDb, dbLogger)
	if err == nil {
		Logger.Infof("app loaded service Db successfully! ")
	}
	return
}
*/
// 加载redis
/*func LoadRedis() (err error) {
	if Redis != nil {
		return nil
	}
	// Redis, err = lib.NewRedis(cfgRedis)
	Redis, err = singleton.NewRedisClient(config.Cfg.Redis.Host, config.Cfg.Redis.Password, config.Cfg.Redis.DbNum, 100)
	if err == nil {
		Logger.Infof("app loaded service Redis successfully! ")
	}
	return
}*/
