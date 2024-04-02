package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	num := 1
	for num <= n {
		fmt.Print(num, " ")
		num *= 2
	}
}
