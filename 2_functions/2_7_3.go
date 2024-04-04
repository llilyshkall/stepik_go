package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var str string
	fmt.Scan(&str)
	fmt.Printf("%c", str[0])
	for i := 1; i < utf8.RuneCountInString(str); i++ {
		fmt.Printf("*%c", str[i])
	}
}
