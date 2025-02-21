package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Not enough arguments. :(")
		return
	}
	var total, nInts, nFloats int
	invalid := make([]string, 0)
	for _, k := range arguments[1:] {
		_, err := strconv.Atoi(k)
		if err == nil {
			total++
			nInts++
			continue
		}
		_, err = strconv.ParseFloat(k, 64)
		if err == nil {
			total++
			nFloats++
			continue
		}
		invalid = append(invalid, k)
	}
	fmt.Printf("Results: Total: %d, Floats: %d, Ints: %d\n", total, nInts, nFloats)
	fmt.Println("Invalid entries:")
	for i, v := range invalid {
		fmt.Printf("%d: \"%s\"\n", i, v)
	}
}
