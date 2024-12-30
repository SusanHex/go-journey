package main

import "fmt"

func minMax(x, y int) (min, max int) {
	if x > y {
		min = y
		max = x
		return min, max
	}
	min = x
	max = y
	return
}

func main() {
	x, y := 20, 30
	fmt.Println(minMax(x, y))
	fmt.Println(minMax(-1, 0))
}
