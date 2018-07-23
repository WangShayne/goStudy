package main

import "fmt"

func testSlice0() {
	var a []int
	if a == nil {
		fmt.Println("a is nil")
	} else {
		fmt.Printf("a = %v\n", a)
	}
}

func testSlice1() {
	var a [5]int
	a = [5]int{1, 2, 3, 4, 5}
	b := a[1:3]
	fmt.Println(b)
}

func testSlice2() {
	a := []int{1, 2, 3}
	fmt.Println(a)
}

func testSlice3() {
	var a [5]int
	a = [5]int{1, 2, 3, 4, 5}
	b := a[1:3]
	fmt.Println(b)
	c := a[1:]
	fmt.Println(c)
	d := a[:5]
	fmt.Println(d)
	e := a[:]
	fmt.Println(e)
}

func testSlice4() {
	var a [5]int
	a = [5]int{1, 2, 3, 4, 5}
	s1 := a[1:3]
	for i := range s1 {
		s1[i] += 10
	}
	fmt.Println(a, s1)
}

func testSlice5() {
	var a [5]int
	a = [5]int{1, 2, 3, 4, 5}
	s1 := a[:]
	s2 := a[:]
	s1[0] = 100
	fmt.Println(a, s2)
	s2[1] = 200
	fmt.Println(a, s1)
}

func main() {
	// testSlice0()
	// testSlice1()
	// testSlice2()
	// testSlice3()
	// testSlice4()
	testSlice5()
}
