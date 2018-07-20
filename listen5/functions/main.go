package main

import (
	"fmt"
)

func add(a int, b int, c ...int) (sum, sub int) {
	for i := 0; i < len(c); i++ {
		a = a + c[i]
	}
	sum = a + b
	sub = a - b
	return
}

func main() {
	sum, sub := add(1, 2, 0, 1)
	fmt.Printf("%d \n %d", sum, sub)
}
