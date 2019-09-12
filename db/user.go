package db

import (
	"github.com/jinzhu/gorm"
	"myiris/common"
	"strconv"
)

type User struct {
	gorm.Model
	Password    string
	Name        string
	PhoneNum    string
	Age         int
	Sex         string
	CommentTime int64
	CommentCnt  int
}

func NewUser(id uint, pwd string, name string, phone string, age int, sex string) *User {
	u := &User{
		Password: com.Md5(pwd),
		Name:     name,
		PhoneNum: phone,
		Age:      age,
		Sex:      sex,
	}
	u.ID = id
	return u
}

func (u *User) save() {
	dbCon.Save(u)
}

func GetUser(id uint) *User {
	u := &User{}
	dbCon.Where("id = ?", strconv.Itoa(int(id))).First(u)
	return u
}

func SignUp(u *User) bool {
	id := GetUser(u.ID).ID
	if id > 0 {
		return false
	}
	if !checkUserData(u) {
		return false
	}
	UpdateUserInfo(u)
	return true
}

func Login(id uint, pwd string) bool {
	u := GetUser(id)
	if u.Password == com.Md5(pwd) {
		return true
	}
	return false
}

func UpdateUserInfo(uInfo *User) {
	uInfo.save()
}

func checkUserData(u *User) bool {
	if u.ID <= 0 {
		return false
	}
	if u.Password == "" {
		return false
	}
	if u.Name == "" {
		return false
	}
	if u.PhoneNum == "" {
		return false
	}
	if u.Sex != "m" && u.Sex != "f" {
		return false
	}
	if u.Age <= 0 {
		return false
	}
	return true
}
