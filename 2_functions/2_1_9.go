package main

import "fmt"

func fibonacci(n int) int {
	var a, b, c = 1, 0, 0
	for i := 1; i < n; i++ {
		c = b
		b = a
		a = b + c
	}
	return a
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(fibonacci(n))
}
