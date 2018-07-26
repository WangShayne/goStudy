package main

import (
	"fmt"

	"github.com/study/listen11/calc"
)

func main() {
	var sum int
	sum = calc.Add(5, 6)
	fmt.Printf("%d", sum)
}
