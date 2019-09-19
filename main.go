package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"math/rand"
	"myiris/common/timer"
	"myiris/db"
)

func main() {
	for id := 9001; id < 9050; id++ {
		pwd := fmt.Sprintf("pwd_%d", id)
		name := fmt.Sprintf("jager_%d", id)
		phone := fmt.Sprintf("1316060%d", id)
		age := rand.Int()%90 + 10
		sex := "m"
		if age%2 == 0 {
			sex = "f"
		}

		newUser := db.Register(pwd, name, phone, age, sex)
		if newUser == nil {
			fmt.Printf("Register failed, id exist or data error!\n")
		} else {
			fmt.Printf("Register successful, wellcome %s\n", newUser.GetName())
		}
	}

	var phoneNum = "13160676597"
	var user *db.User
	u := db.Login(phoneNum, "pwd_9003")
	if u == nil {
		fmt.Printf("Login failed, account or password error!\n")
		return
	} else {
		fmt.Printf("Login successful, wellcome %s!\n", u.GetName())
	}

	cm := db.NewComment(user.ID, "hello world!")
	db.PublicComment(u, cm)
	fmt.Printf("Run end!--dayNo=%d weekNo=%d\n", timer.GetDayNo(), timer.GetWeekNo())
}
