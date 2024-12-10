package main

import (
	"fmt"
	"strings"
	"unicode"
)

var f = fmt.Printf

func main() {
	hello_world := "Hello Beautiful World!"
	fmt.Printf("To Upper %s\n", strings.ToUpper(hello_world))
	fmt.Printf("To Lower %s\n", strings.ToLower(hello_world))
	fmt.Printf("%s\n", strings.ToTitle("tHis will Be A tItle"))
	fmt.Printf("EqualFold: %v\n", strings.EqualFold(hello_world, strings.ToUpper(hello_world)))
	fmt.Printf("Index of 'W': %v\n", strings.Index(hello_world, "W"))
	fmt.Printf("Count of 'e': %v\n", strings.Count(hello_world, "e"))
	fmt.Printf("Repeating: %v\n", strings.Repeat(hello_world, 3))
	fmt.Printf("TrimSpace: %s\n", strings.TrimSpace("\tHello\t\n"))
	fmt.Printf("TrimLeft: %s\n", strings.TrimLeft("\tHello hi", "\t"))
	fmt.Printf("TrimRight: %s\n", strings.TrimRight("hello\thi\t", "\t"))
	fmt.Printf("Prefix: %v\n", strings.HasPrefix(hello_world, "Hello"))
	fmt.Printf("Suffix: %v\n", strings.HasSuffix(hello_world, "!"))
	fmt.Printf("Fields: %d\n", len(strings.Fields(hello_world)))
	fmt.Printf("%s\n", strings.Split(hello_world, " "))
	fmt.Printf("%s\n", strings.Replace(hello_world, "Beautiful", "Cute", 1))
	fmt.Printf("SpringAfter: %s\n", strings.SplitAfter(hello_world, " "))
	trimFunction := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	fmt.Printf("TrimFunc: %s\n", strings.TrimFunc(hello_world, trimFunction))
}
