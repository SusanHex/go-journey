package main

import (
	"fmt"
	"os"
	"regexp"
)

func matchNameSur(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[A-Z][a-z]*$`)
	return re.Match(t)
}

func main() {
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Println("Please provide a string to test against.")
		os.Exit(1)
	}
	potentialSurname := arguments[1]
	fmt.Println(matchNameSur(potentialSurname))
}
