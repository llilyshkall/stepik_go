package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	var str string
	fmt.Scan(&str)
	res := utf8.RuneCountInString(str) >= 5
	for _, elem := range str {
		res = res && (unicode.IsDigit(elem) || unicode.Is(unicode.Latin, elem))
	}
	if res {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")

	}
}
