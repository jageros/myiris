package db

import (
	"github.com/jinzhu/gorm"
	"strconv"
	"time"
)

type User struct {
	gorm.Model
	Password    string
	Name        string
	PhoneNum    string
	Age         int
	Sex         string
	CommentTime time.Time
}

func NewUser(id uint, pwd string, name string, phone string, age int, sex string) *User {
	u := &User{
		Password: pwd,
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
	var u = &User{}
	if dd, ok := dbs.UserData[id]; ok {
		u = dd
	} else {
		dbCon.Where("id = ?", strconv.Itoa(int(id))).First(u)
		if u.ID > 0 {
			dbs.UserData[u.ID] = u
		}
	}
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
	if u.Password == pwd {
		return true
	}
	return false
}

func UpdateUserInfo(uInfo *User) {
	id := uInfo.ID
	if u, ok := dbs.UserData[id]; ok {
		u.Name = uInfo.Name
		u.PhoneNum = uInfo.PhoneNum
		u.Age = uInfo.Age
		u.Sex = uInfo.Sex
	} else {
		dbs.UserData[uInfo.ID] = uInfo
	}
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
