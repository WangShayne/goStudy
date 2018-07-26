package main

import (
	"fmt"
)

type Address struct {
	Province string
	City     string
}

type User struct {
	Username string
	Sex      string
	address  Address
}

func main() {
	user := &User{
		Username: "wang",
		Sex:      "nan",
		address: Address{
			Province: "北京",
			City:     "beijing",
		},
	}
	fmt.Printf("%v\n", user)

	user01 := User{}
	fmt.Printf("%v\n", user01)
}
