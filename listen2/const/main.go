package main

import "fmt"

func main() {
	// const a = 100
	// const b int = 101
	// const c = "ok"

	// const (
	// 	a int = 100
	// 	b
	// 	c int = 200
	// 	d
	// )

	// const (
	// 	a = iota
	// 	b
	// 	c
	// 	d
	// )

	const (
		a = 1 << iota
		b
		c
		d
	)

	fmt.Println(a, b, c, d)
}
