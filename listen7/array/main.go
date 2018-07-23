package main

import (
	"fmt"
)

func sumArray(b [3]int) int {
	sum := 0
	for _, val := range b {
		sum += val
	}
	return sum
}

func returnSum(b [10]int, target int) {
	for i := 0; i < len(b); i++ {
		other := target - b[i]
		for j := i + 1; j < len(b); j++ {
			if b[j] == other {
				fmt.Println(i, j)
			}
		}
	}
}
func main() {
	returnSum([10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, 8)
}
