package main

import "fmt"

func main() {
	aMap := make(map[string]string)
	aMap["123"] = "456"
	aMap["key"] = "A value"
	for key, value := range aMap {
		fmt.Println("Key:", key, "Value:", value)
	}
	for _, v := range aMap {
		fmt.Print(" # ", v)
	}
	fmt.Println()
}
