package main

import "fmt"

func main() {
	var N, cipher int
	var res []int
	fmt.Scan(&N, &cipher)
	for N > 0 {
		if N%10 != cipher {
			res = append(res, N%10)
		}
		N /= 10
	}
	for i := len(res) - 1; i >= 0; i-- {
		fmt.Print(res[i])
	}
}
