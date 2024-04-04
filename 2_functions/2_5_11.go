package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	var str string
	fmt.Scan(&str)
	for i := 0; i < utf8.RuneCountInString(str); i++ {
		if strings.Count(str, str[i:i+1]) == 1 {
			fmt.Printf("%c", str[i])
		}
	}
}
