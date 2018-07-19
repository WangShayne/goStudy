package main

import "fmt"

func main() {
	switch a := 3; a {
	case 1, 2, 3:
		fmt.Println("a=1")
	case 4, 5, 6:
		fmt.Println("a=2")
	default:
		fmt.Println("error")
	}
}
