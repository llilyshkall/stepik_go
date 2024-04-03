package main

import "fmt"

func minimumFromFour() int {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	var ret = d
	if a < b && a < c && a < d {
		ret = a
	}
	if b < a && b < c && b < d {
		ret = b
	}
	if c < b && c < a && c < d {
		ret = c
	}
	return ret
}

func main() {
	minimumFromFour()
}
