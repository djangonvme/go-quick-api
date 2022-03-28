package app

import (
	"context"
	"github.com/go-quick-api/config"
	"github.com/go-quick-api/pkg/singleton"
	"time"
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

func TestInstance() {
	ctx := context.Background()
	Redis().Set(ctx, "testx", "123456", time.Second*10)
	v, err := Redis().Get(ctx, "testx").Result()
	if v != "123456" {
		panic("test redis not pass")
	}
	if err != nil {
		panic("test redis error: " + err.Error())
	}
	LoggerInstance.Infof("test instances ok!")
}
