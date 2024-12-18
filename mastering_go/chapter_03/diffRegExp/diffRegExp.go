package main

import (
	"fmt"
	"regexp"
)

func main() {
	//this is a raw string literal
	var re string = `^.*(?=.{7,})(?=.*\d)$`
	exp1, err := regexp.Compile(re)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("RegExp:", exp1)
	exp2 := regexp.MustCompile(re)
	fmt.Println("exp2:", exp2)
}
