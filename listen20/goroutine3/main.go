package main

import (
	"fmt"
	"time"
)

func ptNumber() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d", i)
	}
}

func ptByte() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c", i)
	}
}

func main() {
	go ptNumber()
	go ptByte()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main stop")
}
