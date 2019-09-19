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
	commentData map[uint]*Comment
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
	if !db.HasTable(&Comment{}) {
		db.CreateTable(&Comment{})
	}
	dbCon = db
	dbs = &DatabaseData{}
	dbs.commentData = map[uint]*Comment{}
	defer db.Close()
}

func Save(ids ...uint) {
	if len(ids) > 0 {
		for _, id := range ids {
			if u, ok := dbs.commentData[id]; ok {
				u.save()
			}
		}
	} else {
		for _, u := range dbs.commentData {
			u.save()
		}
	}
}
