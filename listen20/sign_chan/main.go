// 单向管道

package main

import (
	"fmt"
)

func sendData(sendch chan<- int) {
	sendch <- 1000
}

func readData(sendch <-chan int) {

	data := <-sendch
	fmt.Println(data)
}

func main() {
	var sendch chan int
	sendch = make(chan int)
	go sendData(sendch)
	readData(sendch)
}
