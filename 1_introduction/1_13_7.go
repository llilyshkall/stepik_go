package main

import "fmt"

func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	res := (a + b) / 2
	fmt.Print(res)
}
