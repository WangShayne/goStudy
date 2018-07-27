package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Id   int
	Name string
	Sex  string
}

type Class struct {
	Name    string
	Count   int
	Student []*Student
}

func main() {
	var c Class = Class{
		Name:  "class01",
		Count: 100,
	}
	for i := 0; i < 10; i++ {
		stu := &Student{
			Id:   i,
			Name: fmt.Sprintf("%s%d", "STUDENT", i),
		}
		c.Student = append(c.Student, stu)
	}

	data, _ := json.Marshal(c)
	fmt.Printf("%s", data)
	// cData := json.Unmarshal(data, interface{})
	// fmt.Printf("%v", cData)
}
