package main

import (
	"os"
)

func main() {
	var buff [16]byte
	os.Stdin.Read(buff[:])
	os.Stdout.Write(buff[:])
}
