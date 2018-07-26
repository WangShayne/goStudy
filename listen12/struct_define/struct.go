package main

import (
	"fmt"
)

type User struct {
	Username  string
	Sex       string
	Age       int
	AvatarUrl string
}

func main() {
	var user User
	user.Username = "tom"
	user.Sex = "男"
	user.Age = 15
	user.AvatarUrl = "http://baidu.com/image/xx.jpg"

	fmt.Printf("username = %s\n usersex=%s\n userAge = %d\n userAVA=%s\n", user.Username, user.Sex, user.Age, user.AvatarUrl)

	var user2 User = User{
		Username:  "emma",
		Sex:       "女",
		Age:       19,
		AvatarUrl: "http://www.shayne.wang/image/1.jpg",
	}
	fmt.Printf("%#v", user2)
}
