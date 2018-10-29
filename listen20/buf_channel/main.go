package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	ch <- "hello"
	// ch <- "world"
	// ch <- "llll"
	s1 := <-ch
	// s2 := <-ch
	fmt.Println(s1)
}
