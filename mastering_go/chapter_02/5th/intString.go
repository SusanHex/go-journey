package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 1000
	input := strconv.Itoa(n)
	fmt.Printf("stconv.Itoa output: %v and %T\n", input, input)
	input = strconv.FormatInt(int64(n), 10)
	fmt.Printf("strconv.FormatInt: %v and %T\n", input, input)
	input = string(n)
	fmt.Printf("string(): %v and %T\n", input, input)

}
