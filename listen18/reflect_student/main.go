package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Sex  int
	Age  int
}

func main() {
	var s Student
	v := reflect.ValueOf(&s)
	v.Elem().Field(0).SetString("xiaoming")
	v.Elem().FieldByName("Sex").SetInt(2)
	v.Elem().FieldByName("Age").SetInt(18)
	fmt.Printf("s  = %#v\n", s)
	fmt.Println(s.Name)
}
