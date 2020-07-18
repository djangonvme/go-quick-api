package config

// config.ini 文件的struct
type Config struct {
	General struct {
		Env             string `json:"env"`
		ApiPort         int    `json:"api_port"`
		JwtSecret       string `json:"jwt_secret"`
		TokenExpire     int64  `json:"token_expire"`
		LogDir          string `json:"log_dir"`
		DefaultPageSize uint   `json:"default_page_size"`
		MaxPageSize     uint   `json:"max_page_size"`
	} `json:"general"`

	Redis struct {
		Host     string `json:"host"`
		Password string `json:"password"`
		DbNum    int    `json:"db_num"`
	} `json:"redis"`

	Database struct {
		Schema   string `json:"schema"`
		Host     string `json:"host"`
		User     string `json:"user"`
		Password string `json:"password"`
		Database string `json:"database"`
	} `json:"database"`
}

func (c *Config) Check() error {
	return nil
}
