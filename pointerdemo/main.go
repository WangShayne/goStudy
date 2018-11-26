package main

import (
	"fmt"
)

func main() {
	var i int = 9
	fmt.Println(i)
	fmt.Println("i de di zhi ", &i)

	var ptr *int = &i
	fmt.Println(ptr)
	fmt.Println(&ptr)
	fmt.Println(*ptr)
}
