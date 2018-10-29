package main

import (
	"fmt"
	"reflect"
)

// func test_interface(a interface{}) {
// 	t := reflect.TypeOf(a)
// 	fmt.Printf("type of a is:%v\n", t)

// 	k := t.Kind()
// 	fmt.Printf("%v", k)
// }

func test_reflectVal(a interface{}) {
	t := reflect.ValueOf(a)
	fmt.Println(t.Type())
	v := t.Type()
	fmt.Println(v)
	k := t.Kind()
	switch k {
	case reflect.Float64:
		t.SetFloat(1.4)
		fmt.Printf("a is float32 %d\n", t.Float())
	case reflect.Int:
		fmt.Printf("a is Int32 %d\n", t.Int())
	case reflect.Ptr:
		t.Elem().SetFloat(6.8)
	}

}

func main() {
	var x float64
	x = 1.2
	// test_interface(x)
	test_reflectVal(&x)
	fmt.Println(x)
}
