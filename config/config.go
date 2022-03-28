package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
	"github.com/go-quick-api/pkg/util"
)

func LoadConfig(cf string) func() (*Config, error) {
	return func() (*Config, error) {
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
		fmt.Printf("load config success! config: %v \n", util.ToJson(cg))
		return cg, nil
	}
}

type Config struct {
	General struct {
		Env         string `json:"env"`
		JwtSecret   string `json:"jwtSecret"`
		TokenExpire int64  `json:"tokenExpire"`
		LogDir      string `json:"logDir"`
	} `json:"general"`

	Server struct {
		Addr string `json:"addr"`
	} `json:"server"`

	Redis struct {
		Host     string `json:"host"`
		Password string `json:"password"`
		Db       int    `json:"db"`
	} `json:"redis"`

	MySQL struct {
		Host     string `json:"host"`
		User     string `json:"user"`
		Password string `json:"password"`
		DbName   string `json:"dbname"`
	} `json:"mysql"`
}

func (c *Config) Check() error {
	return nil
}
