package main

import (
	"fmt"
)

func main() {
	var s map[string][]int
	s = make(map[string][]int)
	key := "sot1"
	value, ok := s[key]
	if !ok {
		value = make([]int, 0, 16)
	}
	value = append(value, 100)
	value = append(value, 200)
	value = append(value, 300)
	s[key] = value
	fmt.Println(s)
}
