package main

import (
	"fmt"
)

func main() {
	var c chan int
	fmt.Printf("c is %v :\n", c)
	c = make(chan int, 1)
	c <- 1000
	fmt.Printf("c is %v :\n", c)
	data := <-c
	fmt.Printf("data is %v\n", data)
}
