package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Print(n, " ")
	if n%10 == 1 && n != 11 {
		fmt.Print("korova")
	}
	if n%10 >= 2 && n%10 <= 4 && !(n >= 11 && n <= 14) {
		fmt.Print("korovy")
	}
	if n%10 >= 5 || n%10 == 0 || (n >= 11 && n <= 14) {
		fmt.Print("korov")
	}
}
