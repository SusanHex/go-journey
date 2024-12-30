package main

import (
	"fmt"
	"sort"
)

type Grades struct {
	Name    string
	Surname string
	Grade   int
}

func main() {
	data := []Grades{{"Junior", "G", 10}, {"Susan", "H", 23}, {"T", "B", 34}, {"E", "B", 23}}
	isSorted := sort.SliceIsSorted(data, func(i, j int) bool {
		return data[i].Grade < data[j].Grade
	})
	if isSorted {
		fmt.Println("It is sorted!")
	} else {
		fmt.Println("it is not sorted")
	}
	sort.Slice(data, func(i, j int) bool {
		return data[i].Grade < data[j].Grade
	})
	fmt.Println(data)
}
