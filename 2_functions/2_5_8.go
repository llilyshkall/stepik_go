package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	n := utf8.RuneCountInString(text)
	str := []rune(text)
	res := true
	for i := 0; i < n/2; i++ {
		res = res && str[i] == str[n-1-i]
	}
	if res {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}
