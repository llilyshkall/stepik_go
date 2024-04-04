package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Scan(&str)
	fmt.Printf("%c", MaxCipher(str))
}

func MaxCipher(str string) rune {
	s := []rune(str)
	ret := '0'
	for _, elem := range s {
		if elem > ret {
			ret = elem
		}
	}
	return ret
}
