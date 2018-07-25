package main

import (
	"fmt"
)

func main() {
	var a map[string]int
	fmt.Println(a)
	a = make(map[string]int, 16)
	fmt.Println(a)
	a["stu01"] = 100
	a["stu02"] = 100
	a["stu03"] = 100
	fmt.Println(a)

	b := map[string]int{
		"stu01": 100,
		"stu02": 101,
		"stu03": 102,
	}
	b["stu01"] = 1000
	b["stu04"] = 111
	fmt.Println(b)
}
