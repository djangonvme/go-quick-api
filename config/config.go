package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"gitlab.com/task-dispatcher/pkg/util"
)

func LoadConfig(cf string) func() (*Config, error) {
	return func() (*Config, error) {
		if cf == "" {
			return nil, fmt.Errorf("config file is emtpy")
		}
		if ok, _ := util.IsPathExists(cf); !ok {
			return nil, fmt.Errorf("%s is not exists", cf)
		}
		cg := &Config{}
		_, err := toml.DecodeFile(cf, cg)
		if err != nil {
			return nil, err
		}
		fmt.Printf("load config success! config: %v", util.ToJson(cg))
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
		Host string `json:"host"`
		Port int    `json:"port"`
	} `json:"server"`

	Redis struct {
		Host     string `json:"host"`
		Password string `json:"password"`
		DbNum    int    `json:"dbnum"`
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
