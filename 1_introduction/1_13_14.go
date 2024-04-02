package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	a := 1
	b := 0
	c := 0
	n := 1
	for a < N {
		c = b
		b = a
		a = b + c
		n++
	}
	if a == N {
		fmt.Print(n)
	} else {
		fmt.Print(-1)
	}
}
