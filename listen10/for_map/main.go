package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	var a map[string]int = make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		k := fmt.Sprintf("stu%d", i)
		v := rand.Intn(1000)
		a[k] = v
	}
	for key, value := range a {
		fmt.Printf("map[%s]=%d\n", key, value)
	}

}
