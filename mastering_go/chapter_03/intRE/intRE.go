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
	arguments := os.Args[1:]
	if len(arguments) < 2 {
		fmt.Println("Please provide a string to test against.")
		os.Exit(1)
	}
	var trueMatch, nonMatch uint
	for _, input := range arguments {
		isMatch := matchInt(input)
		if isMatch {
			trueMatch += 1
		} else {
			nonMatch += 1
		}
		fmt.Print(isMatch, " ")
	}
	fmt.Println("\nIs int:", trueMatch, "Is not:", nonMatch)
}
