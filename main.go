package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	com "myiris/common"
	"myiris/db"
)

func main() {
	//for id := 9001; id < 9050; id++ {
	//	pwd := fmt.Sprintf("pwd_%d", id)
	//	name := fmt.Sprintf("jager_%d", id)
	//	phone := fmt.Sprintf("1316060%d", id)
	//	age := rand.Int()%90 + 10
	//	sex := "m"
	//	if age%2 == 0 {
	//		sex = "f"
	//	}
	//	newUser := db.NewUser(uint(id), pwd, name, phone, age, sex)
	//	if db.SignUp(newUser) {
	//		fmt.Printf("Sign up successful!\n")
	//	} else {
	//		fmt.Printf("Sign up failed, id exist or data error!\n")
	//	}
	//}
	var uid = uint(9003)
	if db.Login(uid, "pwd_9003") {
		u := db.GetUser(uid)
		fmt.Printf("Login successful, wellcome %s!\n", u.Name)
	} else {
		fmt.Printf("Login failed, account or password error!\n")
		return
	}
	cm := db.NewComment(uid, "hello world!")
	db.PublicComment(cm)
	fmt.Printf("Run end!--dayNo=%d weekNo=%d\n", com.GetDayNo(), com.GetWeekNo())
}
