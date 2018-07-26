package main

import "fmt"

type User struct {
	UserName string
	Sex      string
	int
	string
}

func main() {
	var user User
	user.UserName = "tom"
	user.Sex = "123"
	user.int = 20
	user.string = "ssss"
	fmt.Printf("%#v", user)
}
