package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	m := n/100 + n/10%10*10 + n%10*100
	fmt.Println(m)
}
