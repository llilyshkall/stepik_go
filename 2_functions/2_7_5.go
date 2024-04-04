package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)
	for _, elem := range []rune(str) {
		c := elem - '0'
		c *= c
		fmt.Print(c)
	}
}
