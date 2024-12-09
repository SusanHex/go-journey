package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide an argument :(")
		return
	}
	argument := os.Args[1]
	value, err := strconv.Atoi(argument)
	if err != nil {
		fmt.Println("Can not convert argument to int: ", argument)
		return
	}
	switch {
	case value == 0:
		fmt.Println("Zero! :)")
	case value > 0:
		fmt.Println("Positive Integer")
	case value < 0:
		fmt.Println("Negative Integer")
	default:
		fmt.Println("Something went horribly wrong. Here is the argument: ", argument)
	}
}
