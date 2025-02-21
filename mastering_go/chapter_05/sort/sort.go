package main

import (
	"fmt"
	"sort"
)

type Size struct {
	F1 int
	F2 string
	F3 int
}

// We want to sort Person records based on the value of Size.F1

type Person struct {
	F1 int
	F2 string
	F3 Size
}

type Personslice []Person

func (a Personslice) Len() int {
	return len(a)
}
func (a Personslice) Less(i, j int) bool {
	return a[i].F3.F3 < a[j].F3.F3
}
func (a Personslice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	data := []Person{
		Person{1, "One", Size{1, "Person_1", 10}},
		Person{2, "Two", Size{2, "Person_2", 20}},
		Person{3, "Three", Size{3, "Person_3", -20}},
	}
	fmt.Println("Before:", data)
	sort.Sort(Personslice(data))
	fmt.Println("After:", data)
	sort.Sort(sort.Reverse(Personslice(data)))
	fmt.Println("Reverse:", data)

}
