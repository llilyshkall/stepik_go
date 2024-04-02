package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for n >= 10 {
		s := 0
		for n > 0 {
			s += n % 10
			n /= 10
		}
		n = s
	}
	fmt.Print(n)
}
