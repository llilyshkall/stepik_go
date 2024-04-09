package main

import "fmt"

func work(a int) int {
	return a * a
}

func main() {
	m := make(map[int]int)
	for range 10 {
		var a int
		fmt.Scan(&a)
		if _, ok := m[a]; !ok {
			m[a] = work(a)
		}
		fmt.Print(m[a], " ")
	}
}
