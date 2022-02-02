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
		return Redis()
	} else {
		dLocker = NewProcessLocker()
		return dLocker
	}
}
