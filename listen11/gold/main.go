package main

import (
	"fmt"

	"github.com/study/listen11/cutGold"
)

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Eilzabeth", "Noo",
	}
	distribution = make(map[string]int, len(users))
)

func main() {
	result, resultMap := cutGold.Cut(coins, users, distribution)
	fmt.Println(result)
	fmt.Printf("%#v", resultMap)
}
