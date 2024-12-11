package main

import "fmt"

func main() {
	a := make([]int, 4)
	fmt.Println("L:", len(a), "C:", cap(a))
	b := []int{0, 1, 2, 3, 4, 5}
	fmt.Println("L:", len(b), "C:", cap(b))
	aSlice := make([]int, 4, 4)
	fmt.Println(aSlice)
	aSlice = append(aSlice, 5)
	fmt.Println(aSlice)
	// Capacity is doubled
	fmt.Println("L:", len(aSlice), "C:", cap(aSlice))
	aSlice = append(aSlice, []int{9, 10, 11, 12, 13, 14}...)
	fmt.Println(aSlice)
	fmt.Println("L:", len(aSlice), "C:", cap(aSlice))
}
