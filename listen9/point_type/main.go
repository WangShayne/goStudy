package main

import "fmt"

func test1() {
	var a int32
	a = 1000
	fmt.Printf("address=%p,%d\n", &a, a)

	var b *int32
	fmt.Printf("address b :%p,%v\n", &b, b)
	b = &a
	*b = 300
	fmt.Println(*b)
}

func test2() {
	var a int = 200
	var b *int = &a
	fmt.Printf("bcun%d", *b)
	*b = 1000
	fmt.Printf("%d\n", *b)
	fmt.Printf("a=%d", a)
}

func test3() {
	a := 58
	fmt.Printf("now a = %d", a)
	b := &a
	modify(b)
	fmt.Printf("now a = %d", a)
}

func modify(val *int) {
	*val = 55
}

func main() {
	// test1()
	// test2()
	test3()
}
