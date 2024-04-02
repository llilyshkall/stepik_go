package main

import "fmt"

func main() {
	var N, min_elem, elem, count int
	fmt.Scan(&N)
	fmt.Scan(&min_elem)
	count = 1
	for i := 1; i < N; i++ {
		fmt.Scan(&elem)
		if elem == min_elem {
			count++
		}
		if elem < min_elem {
			min_elem = elem
			count = 1
		}
	}
	fmt.Print(count)
}
