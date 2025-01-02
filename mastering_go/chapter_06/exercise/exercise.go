package main

import (
	"fmt"
	"sort"
)

func NamedSort(x, y, z int) (min, mid, max int) {
	temp := []int{x, y, z}
	sort.Sort(sort.IntSlice(temp))
	min, mid, max = temp[0], temp[1], temp[2]
	return
}

func NonNamedSort(x, y, z int) (int, int, int) {
	temp := []int{x, y, z}
	sort.Sort(sort.IntSlice(temp))
	min, mid, max := temp[0], temp[1], temp[2]
	return min, mid, max
}

func main() {
	min, mid, max := NamedSort(1, 34, 5)
	fmt.Println(min, mid, max)
	min, mid, max = NonNamedSort(2, 2000, 3)
	fmt.Println(min, mid, max)
}
