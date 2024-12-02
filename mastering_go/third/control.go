package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide an argument :(")
		return
	}
	argument := os.Args[1]
	switch argument {
	case "0":
		fmt.Println("Zero! :)")
	case "1":
		fmt.Println("One!!!!!!!")
	case "2", "3", "4":
		fmt.Println("You entered 2, 3 or 4")
	default:
		fmt.Println("Value: ", argument)
	}
}
