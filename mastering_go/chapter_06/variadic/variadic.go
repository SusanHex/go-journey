package main

import (
	"fmt"
	"os"
)

func addFloats(message string, sFloats ...float64) float64 {
	fmt.Println(message)
	sum := float64(0)
	for _, a := range sFloats {
		sum += a
	}
	return sum
}

func everything(input ...interface{}) {
	fmt.Println(input)
}

func main() {
	sum := addFloats("Adding floats!...", 1.1, 2.3, 2.3, 4.6, 50.3)
	fmt.Println("Sum:", sum)
	s := []float64{1.0, 2.3, 4.5, 5.3}
	sum = addFloats("Adding floats!", s...)
	fmt.Println("Sum:", sum)
	everything(s)
	// Can not directly pass []string as []interface{}
	// It has to be converted
	empty := make([]interface{}, len(os.Args[1:]))
	for i, v := range os.Args[1:] {
		empty[i] = v
	}

	everything(empty...)
	arguments := os.Args[1:]
	empty = make([]interface{}, len(os.Args))
	for i := range arguments {
		empty[i] = arguments[i]
	}
	everything(empty...)

	str := []string{"One", "Two", "Three"}
	everything(str, str, str)
}
