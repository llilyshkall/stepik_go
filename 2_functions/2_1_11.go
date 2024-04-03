package main

import "fmt"

func sumInt(data ...int) (n, sum int) {
	n = len(data)
	sum = 0
	for _, elem := range data {
		sum += elem
	}
	return n, sum
}

func main() {
	a, b := sumInt(1, 0)
	fmt.Println(a, b)
}
