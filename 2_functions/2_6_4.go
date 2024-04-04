package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b != 0 {
		return a / b, nil
	}
	return 0, errors.New("divide by zero")
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	res, err := divide(a, b)
	if err != nil {
		fmt.Println("ошибка")
	} else {
		fmt.Println(res)
	}
}
