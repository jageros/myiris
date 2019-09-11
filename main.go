package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"myiris/conf"
)

type User struct {
	gorm.Model
	Name     string
	PhoneNum string
	Age      int
	Sex      string
}

func main() {
	cfg := conf.GetDBCfg()
	fmt.Printf("cfg=%v\n", cfg)
	url := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	db, err := gorm.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	u := User{Name:"jay", PhoneNum:"13160676598", Age:22, Sex:"female"}
	if !db.HasTable(&u) {
		db.CreateTable(&u)
	}
	db.Create(&u)
	db.NewRecord(u)
	db.Save(&u)
	fmt.Printf("Run end!\n")
}
