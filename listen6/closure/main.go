package main

import (
	"fmt"
	"strings"
	"time"
)

// Adder need int
func Adder() func(int) int {

	var x int
	return func(b int) int {
		x += b
		return x
	}
}

func test() {
	f := Adder()
	val := f(1)
	fmt.Printf("%d\n", val)
	val = f(20)
	fmt.Printf("%d\n", val)
	val = f(100)
	fmt.Printf("%d\n", val)

	f1 := Adder()
	val = f1(1)
	fmt.Printf("%d\n", val)
}

func test2(base int) func(int) int {
	return func(i int) int {
		base += i
		return base
	}
}
func test3(base string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, base) {
			return name + base
		}
		return name
	}
}

func test4(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func test5() {
	for i := 0; i < 5; i++ {
		go func(index int) {
			fmt.Println(index)
		}(i)
	}

	time.Sleep(time.Second)
}

func main() {
	// test()
	// t1 := test2(10)
	// fmt.Println(t1(1), t1(2))
	// t2 := test2(100)
	// fmt.Println(t2(1), t2(2))
	// f1 := test3(".jpeg")
	// f2 := test3(".bmp")
	// fmt.Println(f1("test1.jpeg"))
	// fmt.Println(f2("test2"))
	// f1, f2 := test4(10)
	// fmt.Println(f1(1), f2(2))  // 11 9
	// fmt.Println(f1(3), f2(4))  // 12 8
	// fmt.Println(f1(5), f2(6))  // 13 7
	// fmt.Println(f1(7), f2(8))  // 14 6
	// fmt.Println(f1(9), f2(10)) // 15 5
	test5()
}
