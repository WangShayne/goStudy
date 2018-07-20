package main

import "fmt"

func justIfy(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func example1() {
	for i := 2; i < 100; i++ {
		if justIfy(i) == true {
			fmt.Println(i)
		}
	}
}

func isSXH(n int) bool {

	first := n % 10
	second := n / 10 % 10
	thrid := n / 100 % 10
	sum := first*first*first + second*second*second +thrid*thrid*thrid
	if sum == n {
		return true
	}else {
		return false
	}
}

func example2() {
	for i := 100; i < 1000; i++ {
		if isSXH(i){
			fmt.Printf("%d is SXH",i)
		}
	}
}

func main() {
	// example1()
	example2()
}
