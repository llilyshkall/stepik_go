package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	text, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err == nil || err == io.EOF {
		text = strings.Replace(text, " ", "", -1)
		text = strings.Replace(text, "\n", "", -1)
		text = strings.Replace(text, ",", ".", -1)
		var nums []string = strings.Split(text, ";")
		first, _ := strconv.ParseFloat(nums[0], 64)
		second, _ := strconv.ParseFloat(nums[1], 64)
		fmt.Printf("%.4f", first/second)
	}
}
