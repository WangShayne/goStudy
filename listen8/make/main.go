package main

import "fmt"

func testMake() {
	var a []int
	a = make([]int, 5, 5)
	a[0] = 10
	a[1] = 20
	fmt.Println(a)
}

func testMake1() {
	var a []int
	a = make([]int, 1, 5)
	a[0] = 10
	// a[1] = 20
	a = append(a, 1)
	fmt.Println(a)
}

func testCap() {
	a := [...]int{1, 2, 3, 4, 5}
	b := a[1:3]
	fmt.Println(b)
	b = append(b, 1)
	b = append(b, 1)
	b = append(b, 1)
	b = append(b, 1)
	fmt.Println(len(b), cap(b))
	fmt.Println(a, b)

}

func main() {
	// testMake()
	// testMake1()
	testCap()
}
