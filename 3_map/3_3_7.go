package main

import (
	"fmt"
	"strconv"
)

func main() {
	fn := func(a uint) uint {
		var ret uint
		var s string = strconv.FormatUint(uint64(a), 10)
		for _, r := range []rune(s) {
			if r != '0' && (r-'0')%2 == 0 {
				ret = ret*10 + uint(r-'0')
			}
		}
		if ret == 0 {
			ret = 100
		}
		return ret
	}
	fmt.Println(fn(11120))
}
