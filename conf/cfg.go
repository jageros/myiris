package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

var cfg *Cfg

type Cfg struct {
	DBConfig   *DatabaseCfg `toml:"database"`
	CommentCfg *Comment     `toml:"comment"`
}

type DatabaseCfg struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	DBName   string `toml:"dbName"`
}

type Comment struct {
	Cnt int `toml:"cnt"`
}

func init() {
	var config *Cfg
	if _, err := toml.DecodeFile("conf/config.toml", &config); err != nil {
		fmt.Println(err)
	}
	cfg = config
}

func GetDBCfg() *DatabaseCfg {
	return cfg.DBConfig
}

func GetCommentCfg() *Comment {
	return cfg.CommentCfg
}
