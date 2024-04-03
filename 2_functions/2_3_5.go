package main

import "fmt"

func test(x1 *int, x2 *int) {
	fmt.Println(*x1 * *x2)
}
func main() {
	var x1, x2 = 2, 2
	test(&x1, &x2)
}
