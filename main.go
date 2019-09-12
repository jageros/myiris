package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"math/rand"
	"myiris/db"
)



func main() {
	u := db.GetData(6)
	fmt.Printf("user=%+v\n", u)
	for id := 8001; id < 8050; id++ {
		pwd := fmt.Sprintf("pwd_%d", id)
		name := fmt.Sprintf("jager_%d", id)
		phone := fmt.Sprintf("1316060%d", id)
		age := rand.Int()%90+10
		sex := "m"
		if age%2 == 0 {
			sex = "f"
		}
		newUser := db.NewUser(uint(id), pwd, name, phone, age, sex)
		if db.SignUp(newUser) {
			fmt.Printf("Sign up successful!\n")
		}else {
			fmt.Printf("Sign up failed, id exist or data error!\n")
		}
	}

	if db.Login(80080, "asdf") {
		u := db.GetData(80080)
		fmt.Printf("Login successful, wellcome %s!", u.Name)
	}else {
		fmt.Printf("Login failed, account or password error!")
	}
	fmt.Printf("\nRun end!\n")
}
