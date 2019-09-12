package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"myiris/conf"
)

var (
	dbCon *gorm.DB
	dbs   *DatabaseData
)

type DatabaseData struct {
	UserData map[uint]*User
}

func init() {
	cfg := conf.GetDBCfg()
	fmt.Printf("Database init Generate db url from cfg=%+v\n", cfg)
	url := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	db, err := gorm.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	if !db.HasTable(&User{}) {
		db.CreateTable(&User{})
	}
	dbCon = db
	dbs = &DatabaseData{}
	dbs.UserData = map[uint]*User{}
}

func Save(ids ...uint) {
	if len(ids) > 0 {
		for _, id := range ids {
			if u, ok := dbs.UserData[id]; ok {
				u.save()
			}
		}
	} else {
		for _, u := range dbs.UserData {
			u.save()
		}
	}
}
