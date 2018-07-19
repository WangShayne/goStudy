package main

import "fmt"
import "time"

func getNow() {
	now := time.Now()
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()

	// fmt格式化
	fmt.Printf("%02d/%02d/%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

	// time格式化
	nowStr := now.Format("2006/01/02 15:04:00")
	fmt.Println(nowStr)
}

func doSomething() {
	fmt.Println("do something")
}

func timeOut() {
	ticker := time.Tick(3 * time.Second)
	for i := range ticker {
		fmt.Printf("di %v ci ", i)
		doSomething()
	}
}

func timeTotal() {
	start := time.Now().UnixNano()
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond)
	}
	end := time.Now().UnixNano()
	fmt.Println((end - start) / 1000)
}

func main() {
	// getNow()
	// timeOut()
	timeTotal()
}
