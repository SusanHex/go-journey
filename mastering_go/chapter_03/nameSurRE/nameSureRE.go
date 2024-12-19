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
	arguments := os.Args[1:]
	if len(arguments) < 2 {
		fmt.Println("Please provide a string to test against.")
		os.Exit(1)
	}
	var falseSurname, trueSurname uint
	for _, potentialSurname := range arguments {
		isSurname := matchNameSur(potentialSurname)
		fmt.Println(isSurname)
		if isSurname {
			trueSurname += 1
		} else {
			falseSurname += 1
		}
	}
	fmt.Println("True:", trueSurname, "False:", falseSurname)
}
