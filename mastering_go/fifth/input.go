package main

import (
	"fmt"
)

func main() {
	// Get User Input
	fmt.Println("What is your name?: ")
	var name string
	fmt.Scanln(&name)
	fmt.Println("I hope you have a wonder day,", name)
}
