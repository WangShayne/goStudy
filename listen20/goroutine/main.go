package main

import (
	"time"
	"fmt"
)

func hello(){
	fmt.Println(" hello this is go goroutine")
}

func main(){
	go hello()
	fmt.Println("main runtime")
	time.Sleep(time.Second)
}
