package db

import (
	"github.com/jinzhu/gorm"
	"github.com/lhj168os/myiris/common/timer"
	"github.com/lhj168os/myiris/common/utils"
	"github.com/lhj168os/myiris/conf"
	"time"
)

type iUser interface {
	OnComment()
	Login(password string) bool
	UpdateScore(score int)
	GetUid() uint
	GetName() string
	CanComment() bool
}

// ==================================user=================================

type User struct {
	gorm.Model
	Password    string
	Name        string
	PhoneNum    string `gorm:"type:varchar(128);not null;index:phone_num"`
	Age         int    `gorm:"default:18"`
	Sex         string
	CommentTime int64
	CommentCnt  int
	Score       int
	VipTime     int64
}

// =============== interface implement ================

func (u *User) OnComment() {
	u.CommentTime = time.Now().Unix()
	u.CommentCnt += 1
	u.save()
}

func (u *User) Login(password string) bool {
	return u.Password == utils.Md5(password)
}

func (u *User) GetUid() uint {
	return u.ID
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) UpdateScore(score int) {
	u.Score += score
	u.save()
}

func (u *User) CanComment() bool {
	if timer.GetDayNo(u.CommentTime) == timer.GetDayNo() {
		cnt := conf.GetCommentCfg().Cnt
		if u.CommentCnt >= cnt {
			return false
		}
	} else {
		u.CommentCnt = 0
		u.save()
	}
	return true
}

// ===================================================

// ================= Get interface ===================

func GetUser(phone string) iUser {
	u := &User{}
	dbCon.Where(&User{PhoneNum: phone}).First(u)
	if u.ID <= 0 {
		return nil
	}
	return u
}

//=====================================================

func (u *User) save() {
	dbCon.Save(u)
}

func (u *User) checkUserData() bool {
	if len(u.Password) < 6 || len(u.PhoneNum) < 10 || len(u.Name) < 3 {
		return false
	}
	if u.Sex != "f" && u.Sex != "m" {
		u.Sex = "m"
	}
	if u.Age <= 0 || u.Age > 200 {
		u.Age = 18
	}
	return true
}

func Register(pwd string, name string, phone string, age int, sex string) iUser {
	user := GetUser(phone)
	if user != nil {
		return nil
	}
	u := &User{
		Password: utils.Md5(pwd),
		Name:     name,
		PhoneNum: phone,
		Age:      age,
		Sex:      sex,
		VipTime:  timer.TimeBaseUnix,
	}
	if !u.checkUserData() {
		return nil
	}
	dbCon.Create(u)
	return u
}

func Login(phone string, pwd string) iUser {
	u := GetUser(phone)
	if u.Login(pwd) {
		return u
	}
	return nil
}

func UpdateUserInfo(uInfo *User) {
	uInfo.save()
}
