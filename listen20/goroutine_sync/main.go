package main

import (
	"fmt"
	"time"
)

func hello(exchan chan bool) {
	time.Sleep(time.Second * 5)
	fmt.Println("hello this is goroutine")
	exchan <- true

}

func main() {
	var exchan chan bool
	exchan = make(chan bool)
	go hello(exchan)
	fmt.Println("main complate")
	<-exchan
}
