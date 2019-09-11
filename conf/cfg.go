package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type Cfg struct {
	DBConfig *DatabaseCfg `toml:"database"`
}

type DatabaseCfg struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	DBName   string `toml:"dbName"`
}

func GetDBCfg() *DatabaseCfg {
	var config *Cfg
	if _, err := toml.DecodeFile("conf/config.toml", &config); err != nil {
		fmt.Println(err)
	}
	return config.DBConfig
}
