package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var str string
	// fmt.Scanf("%s", &str)
	// fmt.Println("reader frm fmt:", str)

	reader := bufio.NewReader(os.Stdin)
	str, _ = reader.ReadString('\n')
	fmt.Println("from bufio:", str)
}
