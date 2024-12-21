package main

import "fmt"

func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func main() {
	strSlice := []string{"Hello", "Beautiful", "World"}
	intSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	floatSlice := []float64{1.2, 1.3, 4.5, 3.6}
	PrintSlice(strSlice)
	PrintSlice(intSlice)
	PrintSlice(floatSlice)
}
