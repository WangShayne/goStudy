package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./file.go")
	if err != nil {
		fmt.Println("open file failde , err:", err)
	}
	defer file.Close()
	var (
		content []byte
		buf     [128]byte
	)
	for {
		n, err := file.Read(buf[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read file:", err)
			return
		}
		content = append(content, buf[:n]...)
	}
	fmt.Println("data", string(content))
}
