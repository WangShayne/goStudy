package main

import (
	"fmt"
)

type MethodUtils struct {
}

func (mu *MethodUtils) printmula(a int) {
	for i := 1; i <= a; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", i, j, i*j)
		}
		fmt.Println()
	}
}

func main() {
	mu := MethodUtils{}
	mu.printmula(9)
}
