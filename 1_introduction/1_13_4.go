package main

import "fmt"

func main() {
	var s int
	fmt.Scan(&s)
	m := s / 60 % 60
	h := s / 3600
	fmt.Println("It is", h, "hours", m, "minutes.")
}
