package main

import "fmt"

func reserveString() {
	var str = "hello"
	fmt.Printf("old str =%s\n", str)
	var byteSlice []byte
	byteSlice = []byte(str)
	for i := 0; i < len(byteSlice)/2; i++ {
		temp := byteSlice[len(byteSlice)-i-1]
		byteSlice[len(byteSlice)-i-1] = byteSlice[i]
		byteSlice[i] = temp
	}
	str = string(byteSlice)
	fmt.Println(str)
}

func main() {
	reserveString()
}
