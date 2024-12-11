package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	b := make([]byte, 12)
	fmt.Println(b)
	b = []byte("Byte Slice おはよう")
	fmt.Println("Byte Slice:", b)
	fmt.Printf("Byte slice as text: %s\n", b)
	fmt.Println("Byte slice as text:", string(b))
	fmt.Printf("Bytes: %d, Runes: %d", len(b), utf8.RuneCountInString(string(b)))
}
