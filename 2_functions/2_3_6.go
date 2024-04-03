package main

import "fmt"

func test1(x1 *int, x2 *int) {
	tmp := *x1
	*x1 = *x2
	*x2 = tmp
	fmt.Println(*x1, *x2)
}

func main() {
	x1, x2 := 2, 4
	test1(&x1, &x2)
}
