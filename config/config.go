package config

import (
	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
	"gitlab.com/qubic-pool/pkg/util"
)

var Instance *Config

func LoadConfig(cf string) (*Config, error) {
	if Instance != nil {
		return Instance, nil
	}
	if cf == "" {
		return nil, errors.Errorf("config file is emtpy")
	}
	if ok, _ := util.IsPathExists(cf); !ok {
		return nil, errors.Errorf("%s is not exists", cf)
	}
	cg := &Config{}
	_, err := toml.DecodeFile(cf, cg)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("load config success! config: %v \n", util.ToJson(cg))
	Instance = cg
	return cg, nil
}

type Config struct {
	Env struct {
		Env string `json:"env"`
	} `json:"env"`

	Jwt struct {
		Expire int64 `json:"expire"`
	} `json:"jwt"`
	Server struct {
		Listen string `json:"listen"`
	} `json:"server"`

	Redis struct {
		Host     string `json:"host"`
		Password string `json:"password"`
		Db       int    `json:"db"`
		PoolSize int    `json:"poolSize"`
	} `json:"redis"`

	MySQL struct {
		Host     string `json:"host"`
		User     string `json:"user"`
		Password string `json:"password"`
		DbName   string `json:"dbname"`
	} `json:"mysql"`

	Log struct {
		Dir string `json:"dir"`
	} `json:"log"`
}

func IsEnvLocal() bool {
	if Instance != nil {
		return Instance.Env.Env == "local"
	}
	return false
}
func IsEnvDebug() bool {
	if Instance != nil {
		return Instance.Env.Env == "debug"
	}
	return false
}
func IsEnvDev() bool {
	if Instance != nil {
		return Instance.Env.Env == "dev"
	}
	return false
}

func IsEnvProduction() bool {
	if Instance != nil {
		return Instance.Env.Env == "production"
	}
	return false
}
