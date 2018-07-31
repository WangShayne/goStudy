package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		a int
		b string
		c float32
	)
	fmt.Fscanf(os.Stdin, "%d%s%f", &a, &b, &c)
	fmt.Println(a, b, c)
}
