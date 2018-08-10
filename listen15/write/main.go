package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	defer file.Close()
	file.Write([]byte("hello world"))
	file.WriteString("string-hello world")
}
