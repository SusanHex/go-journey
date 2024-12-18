package main

import (
	"fmt"
	"os"
	"regexp"
)

func matchInt(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[-+]?\d+$`)
	return re.Match(t)
}

func main() {
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Println("Please provide a string to test against.")
		os.Exit(1)
	}
	input := arguments[1]
	fmt.Println(matchInt(input))
}
