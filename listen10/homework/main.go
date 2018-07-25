package main

import (
	"fmt"
	"strings"
)

func homework1(test map[string]int, str string) {

	// m := map[string]int{
	// 	"how": 0,
	// 	"do":  0,
	// 	"you": 0,
	// }
	for key, value := range test {
		value = strings.Count(str, key)
		fmt.Printf("%s:%d\n", key, value)
	}

}

func homework11(str string) map[string]int {
	var result map[string]int = make(map[string]int, 16)
	words := strings.Split(str, " ")
	for _, value := range words {
		result[value]++
	}
	return result
}

func homework2() {

}

func main() {
	m := map[string]int{
		"how": 0,
		"do":  0,
		"you": 0,
		" ":   0,
	}

	homework1(m, "how do you do")
	result := homework11("how do you do? do you like do something to do?")
	fmt.Println(result)
	// homework2()
}
