// 关闭管道
package main

import (
	"fmt"
)

func producer(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	fmt.Printf("%v", c)
	close(c)
}

func main() {
	var c chan int
	c = make(chan int)
	go producer(c)
	for {
		v, ok := <-c
		if ok == false {
			fmt.Println("channel is closed")
			break
		}
		fmt.Printf("v is :%v,\n ok is :%v\n", v, ok)
	}
}
