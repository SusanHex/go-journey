package main

import (
	"fmt"
	"sort"
)

type Numbers struct {
	value int
}

type NumbersFloat struct {
	value float64
}

type numSlice []Numbers

func (n numSlice) Len() int {
	return len(n)
}

func (n numSlice) Less(i, j int) bool {
	return n[i].value < n[j].value
}

func (n numSlice) Swap(i, j int) {
	temp := n[i].value
	n[i].value = n[j].value
	n[j].value = temp
}

func Print(s interface{}) {
	switch T := s.(type) {
	case NumbersFloat:
		fmt.Println("NumbersFloat type!")
	case Numbers:
		fmt.Println("Numbers type!")
	default:
		fmt.Printf("Unknown type: %T\n", T)

	}
}

func main() {
	numberSlice := numSlice{}
	for i := 10; i > 0; i-- {
		temp := Numbers{value: i}
		numberSlice = append(numberSlice, temp)
	}
	fmt.Println(numberSlice)
	sort.Sort(numberSlice)
	fmt.Println(numberSlice)
	Print(numberSlice)
	Print(NumbersFloat{value: 1.2})
	Print(numberSlice[1])
}
