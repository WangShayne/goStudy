package main

import (
	"fmt"
)

func test1(a int) {
	fmt.Printf("%p\n", &a)
}

func test2(a *int) {
	*a = 1000
}

func main() {
	// test1()
	a := 10
	test2(&a)
	fmt.Println(a)
}
