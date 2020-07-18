package app

import (
	"fmt"
)

const (
	EnvLocal      = "local"
	EnvDev        = "dev"
	EnvTest       = "test"
	EnvProduction = "production"
	CtxStartTime  = "ctx-start-time"
)

type CheckIF interface {
	Check() error
}

// api 监听地址
func HttpServeAddr() string {
	if Cfg != nil {
		return fmt.Sprintf(":%d", Cfg.General.ApiPort)
	}
	return ""
}

// 环境
func IsEnvLocal() bool {
	return CurrentEnv() == EnvLocal
}

func IsEnvDev() bool {
	return CurrentEnv() == EnvDev
}

func IsEnvTest() bool {
	return CurrentEnv() == EnvTest
}

func IsEnvProduction() bool {
	return CurrentEnv() == EnvProduction
}

func CurrentEnv() string {
	if Cfg != nil {
		return Cfg.General.Env
	}
	return ""
}
