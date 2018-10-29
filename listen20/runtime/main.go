package main

import (
	"fmt"
	"runtime"
)

func main() {
	cpu := runtime.NumCPU()
	// 默认所有cpu资源 GOMAXPROCS可选使用的CPU数量
	runtime.GOMAXPROCS(1)
	fmt.Println(cpu)
}
