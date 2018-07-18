package main

import (
	"fmt"
)

func main() {
	var str = "hello"
	fmt.Printf("str[0] = %c,length = %d\n", str[0], len(str))

	for index, val := range str {
		fmt.Printf("str[%d]=%c\n", index, val)
	}

	// 替换第一个字符为x 需转为byte切片
	var byteSlice []byte
	byteSlice = []byte(str)
	byteSlice[0] = 'x'
	str = string(byteSlice)
	fmt.Printf("after motiyf=%s\n", str)

	var str1 = "abc中英文"
	fmt.Printf("str1的bytelength = %d\n", len(str1))
	var runeSlice []rune
	runeSlice = []rune(str1)
	fmt.Printf("str1的字符长度为 = %d\n", len(runeSlice))
}
