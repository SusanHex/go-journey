package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Dang you! I need an argument! Nooooo!!!!")
		return
	}
	var min, max float64
	var initalized = 0
	for i := 1; i < len(arguments); i++ {
		n, err := strconv.ParseFloat(arguments[i], 64)
		if err != nil {
			fmt.Printf("Failed to convert \"%s\" to float. Error: \"%s\"\n", arguments[i], err)
			continue
		}
		if initalized == 0 {
			min = n
			max = n
			initalized = 1
			continue
		}
		if n < min {
			min = n
		} else if n > max {
			max = n
		}
	}
	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}
