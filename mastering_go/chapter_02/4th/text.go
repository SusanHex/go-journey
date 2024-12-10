package main

import "fmt"

func main() {
	aString := "おはようございます"
	fmt.Println("First byte", string(aString[0]))
	r := 'ざ'
	//convert rune to text
	fmt.Printf("As a string: %s and as a character: %c\n", r, r)
	fmt.Println("As an int32 value\n", r)
	fmt.Printf("As a string: %s\n", r)
	fmt.Println()
	for _, v := range aString {
		fmt.Printf("%X ", v)
	}
	fmt.Println()
	for _, v := range aString {
		fmt.Printf(" %c %T", v, v)
	}
	fmt.Println()
}
