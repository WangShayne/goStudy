package main

import (
	"fmt"
	"time"
)

func produce(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
		time.Sleep(time.Second * 2)
	}
	close(c)
}

func main() {
	c := make(chan int)
	go produce(c)
	for v := range c {
		fmt.Printf("reciver:%v\n", v)
	}
}
