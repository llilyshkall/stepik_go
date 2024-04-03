package main

import "fmt"

func vote(x int, y int, z int) int {
	if x+y+z >= 2 {
		return 1
	}
	return 0
}

func main() {
	fmt.Println(vote(1, 0, 1))
}
