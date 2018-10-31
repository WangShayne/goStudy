package main

import (
	routers "github.com/study/gin_test/routers"
)

func main() {
	routers := routers.InitRouter()
	routers.Run(":6789")
}
