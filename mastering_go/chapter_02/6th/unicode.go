package main

import (
	"fmt"
	"unicode"
)

func main() {
	rune_string := "\x99\x00ab\x50\x00\x23\x50\x29\x9c"
	for i := 0; i < len(rune_string); i++ {
		if unicode.IsPrint(rune(rune_string[i])) {
			fmt.Printf("%c\n", rune_string[i])
		} else {
			fmt.Println("Not Printable: ", i)
		}
	}
}
