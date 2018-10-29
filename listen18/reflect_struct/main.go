package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string
	Sex   int
	Age   int
	Score float32
	// xxx   int
}

func main() {
	var s Student
	// a = {
	// 	Name:"xiaoming",
	// 	Sex:0,
	// 	Age:19,
	// }

	v := reflect.ValueOf(s)
	t := v.Type()
	kind := t.Kind()

	switch kind {
	case reflect.Int64:
		fmt.Printf("s is int\n")
	case reflect.Float64:
		fmt.Printf("s is float\n")
	case reflect.Struct:
		fmt.Printf("s is struct\n")
		fmt.Printf("num filed of s is %d \n", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			filed := v.Field(i)
			fmt.Printf("name=%s type=%v value=%v\n",
				t.Field(i).Name, filed.Type(), filed.Interface())
		}
	default:
		fmt.Printf("defautl\n")
	}

}
