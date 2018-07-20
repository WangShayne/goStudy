package main

import "fmt"

func testDefer() {
	defer fmt.Println("defer1")
	fmt.Println("aaa")
	fmt.Println("bbb")
}

func testDefer1() {
	defer fmt.Println("defer1")
	defer fmt.Println("defer2")
	defer fmt.Println("defer3")
	fmt.Println("aaa")
	fmt.Println("bbb")
}

func main() {
	// testDefer()
	testDefer1()
}
