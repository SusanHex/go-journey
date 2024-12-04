package main

import "fmt"

func main() {
	fmt.Println("Conventional Loop")
	for i := 0; i < 10; i++ {
		fmt.Print(i*i, " ")
	}
	fmt.Println()
	fmt.Println("Indiomatic go loop")

	i := 0
	for ok := true; ok; ok = (i != 10) {
		fmt.Print(i*i, " ")
		i++
	}
	fmt.Println()
	fmt.Println("Infinite For Loop with break")
	i = 0
	for {
		if i == 10 {
			break
		}
		fmt.Print(i*i, " ")
		i++
	}
	fmt.Println()
	// Slice example
	fmt.Println("Slice Reading")
	slice_variable := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	for i, v := range slice_variable {
		fmt.Println("Index: ", i, ",", "Value: ", v)
	}
}
