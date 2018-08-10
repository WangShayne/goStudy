package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	str := "hello world"
	err := ioutil.WriteFile("./test.txt", []byte(str), 0755)
	if err != nil {
		fmt.Println(err)
	}
}
