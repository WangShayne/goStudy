package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	content, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(content))
}
