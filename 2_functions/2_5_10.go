package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var str string
	fmt.Scan(&str)
	for i := 1; i < utf8.RuneCountInString(str); i += 2 {
		fmt.Printf("%c", str[i])
	}
}
