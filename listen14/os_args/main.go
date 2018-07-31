package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%s", os.Args[0])
	for index, v := range os.Args {
		if index == 0 {
			continue
		}
		fmt.Printf("os.Args[%d] =%v\n", index, v)
	}
}
