package main

import "fmt"
import "time"

func cale() {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("time ", i, "times")
	}
	fmt.Println("cale close")
}

/*
	同步
	func main() {
		cale()
		fmt.Println("main close")
	}
*/

// 异步
func main() {
	go cale()
	fmt.Println("main close")
	time.Sleep(11 * time.Second)
}
