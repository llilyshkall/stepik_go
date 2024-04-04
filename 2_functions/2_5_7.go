package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	str := []rune(text)
	if unicode.IsUpper(str[0]) && strings.HasSuffix(string(str), ".") {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}
