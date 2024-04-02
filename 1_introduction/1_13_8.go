package main

import "fmt"

func main() {
	var N int
	count := 0
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		var num int
		fmt.Scan(&num)
		if num == 0 {
			count++
		}
	}
	fmt.Print(count)
}
