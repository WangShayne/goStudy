package main

import (
	"fmt"
)

func main() {
	/*
		var a int
		var b bool
		var c string
		var d float32
	*/

	var (
		a int
		b bool
		c string
		d float32
	)
	fmt.Printf("a=%d ,b=%t, c=%s, d=%f\n", a, b, c, d)

	a = 100
	b = true
	c = "ok"
	d = 0.1
	fmt.Printf("a=%d ,b=%t, c=%s, d=%f\n", a, b, c, d)

	/*
		整数 %d
		布尔值 %t
		字符串 %s
		浮点数 %f
	*/

}
