package main

import (
	"fmt"
)

type MethodsUtils struct {
}

func (mu *MethodsUtils) rotate(slc [][]int) [][]int {
	var newSlc = make([][]int, 3)

	for key, val := range slc {
		fmt.Println(val)
		newSlc[key] = make([]int, 3)
		for key1 := range val {
			newSlc[key][key1] = slc[key1][key]
		}
	}

	fmt.Println()

	for _, val := range newSlc {
		fmt.Println(val)
	}

	return nil
}

func main() {
	mu := MethodsUtils{}
	s := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	mu.rotate(s)
}
