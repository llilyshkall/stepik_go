package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	res := b
	for ; res > a && res%7 != 0; res-- {
	}
	if res%7 == 0 {
		fmt.Print(res)
	} else {
		fmt.Print("NO")
	}
}
