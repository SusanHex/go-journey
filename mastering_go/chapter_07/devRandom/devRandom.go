package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("/dev/random")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var seed int64
	binary.Read(f, binary.LittleEndian, &seed)
	fmt.Println("Seed:", seed)
}
