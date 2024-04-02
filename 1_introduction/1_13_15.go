package main

import "fmt"

func main() {
	var N int
	var bin []int
	fmt.Scan(&N)
	for N > 0 {
		bin = append(bin, N%2)
		N /= 2
	}
	for i := len(bin) - 1; i >= 0; i-- {
		fmt.Print(bin[i])
	}
}
