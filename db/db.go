package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopuppy/db"
	"myiris/conf"
)

var DB *gorm.DB

type DBData struct {
	UserData map[int]*User
}

type User struct {
	gorm.Model
	Name     string
	PhoneNum string
	Age      int
	Sex      string
}

func init() {
	cfg := conf.GetDBCfg()
	fmt.Printf("Database init Generate db url from cfg=%v\n", cfg)
	url := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	db, err := gorm.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	if !db.HasTable(&User{}) {
		db.CreateTable(&User{})
	}
	DB = &DBData{}
	DB.UserData = map[int]*User{}
	u := db.Find(&User{})

}

func (d *DBData)GetData() *User {

	u := User{Name:"jay", PhoneNum:"13160676598", Age:22, Sex:"female"}

	db.Create(&u)
	db.NewRecord(u)
	db.Save(&u)
	fmt.Printf("Run end!\n")
}
