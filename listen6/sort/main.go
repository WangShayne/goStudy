package main

import "fmt"

/*
   inputSort 插入排序
*/

func inputSort(a [8]int) [8]int {
	for i := 1; i < len(a); i++ {
		for j := i; j > 0; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			} else {
				break
			}
		}
	}
	return a
}

// selectSort

func selectSort(a [8]int) [8]int {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[i] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	return a
}

// bubbleSort

func bubbleSort(a [8]int) [8]int {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return a
}

func main() {
	i := [8]int{8, 3, 2, 4, 5, 1, 0, 9}
	// inputSort(i)
	// fmt.Println(inputSort(i))
	// fmt.Println(selectSort(i))
	fmt.Println(bubbleSort(i))
}
