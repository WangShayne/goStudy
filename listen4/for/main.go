package main

import (
	"fmt"
)

func forTest1() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", i)
	}
}

func forTest2() {
	for i := 0; i < 10; i++ {
		if i == 2 {
			continue // 跳过当前
		}
		if i > 5 {
			break // 结束循环
		}
		fmt.Printf("%d\n", i)
	}
}

func forTest3() {
	for i, j := 0, 10; i < 10 && j >= 0; i, j = i+1, j-1 {
		fmt.Printf("%d * %d = %d\n", i, j, i*j)
	}
}

func main() {
	// forTest1()
	// forTest2()
	forTest3()
}
