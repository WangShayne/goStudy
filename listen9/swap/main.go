package main

import (
	"fmt"
)

func swap(a *int, b *int) {
	*a, *b = *b, *a
}

func main() {
	a := 100
	b := 333
	swap(&a, &b)
	fmt.Printf("after swap ---> a=%d,b=%d", a, b)
}
