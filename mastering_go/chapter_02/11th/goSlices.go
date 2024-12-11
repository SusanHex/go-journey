package main

import "fmt"

func main() {
	//create an empty slice
	aSlice := []float64{}
	// both len and capacity are 0
	fmt.Println(aSlice, len(aSlice), cap(aSlice))
	//add elements
	aSlice = append(aSlice, 1.12)
	aSlice = append(aSlice, 44.55)
	fmt.Println(aSlice, len(aSlice), cap(aSlice))
	t := make([]int, 4)
	t[0] = -1
	t[1] = -2
	t[2] = -3
	t[3] = -4
	t = append(t, 5)
	fmt.Println(t)
	twoD := [][]int{{1, 2, 3}, {-1, -2, -3}}
	for _, i := range twoD {
		for _, k := range i {
			fmt.Print(k, " ")
		}
		fmt.Println()
	}
	make2D := make([][]int, 2)
	make2D[0] = []int{1, 2, 3}
	make2D[1] = []int{-1, -2, -61}
	fmt.Println(make2D)
}
