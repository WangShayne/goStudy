package main

import "fmt"

func main() {
	const (
		A = iota
		B
		C
		D = 8
		E = 9    //4 iota 自增1
		F = iota //5
		G
	)

	const (
		A1 = iota
		A2
	)

	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
	fmt.Println(G)
	fmt.Println(A1)
	fmt.Println(A2)
}
