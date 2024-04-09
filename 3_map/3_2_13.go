package main

import (
	"fmt"
	"unicode"
)

func strToInt64(s string) int64 {
	var ret int64 = 0
	str := []rune(s)
	for _, value := range str {
		if unicode.IsDigit(value) {
			ret = ret*10 + int64(value) - int64('0')
		}
	}
	return ret
}

func adding(a, b string) int64 {
	return strToInt64(a) + strToInt64(b)
}

func main() {
	var a, b = "absfj23", "dsjk32"
	fmt.Println(adding(a, b))
}
