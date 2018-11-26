package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	ClassName string `json:"class_name"`
}

func main() {

	stu1 := Student{}
	stu1.Name = "shayne wang"
	stu1.Age = 10
	stu1.ClassName = "一年级一班"

	jsonStu1, err := json.Marshal(stu1)
	if err != nil {
		fmt.Println("jsonMarshal error")
	}
	fmt.Println(jsonStu1)
	fmt.Println(string(jsonStu1))
}
