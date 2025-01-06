package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func readSize(f *os.File, size int) []byte {
	buffer := make([]byte, size)
	n, err := f.Read(buffer)

	if err == io.EOF {
		return nil
	} else if err != nil {
		fmt.Println(err)
	}
	return buffer[0:n]
}

func main() {

	filename := flag.String("file", "/etc/hosts", "Please provide a file")
	read_amount := flag.Int("size", 64, "Please provide an amount of bytes to read")
	flag.Parse()
	file, err := os.Open(*filename)
	if err != nil {
		fmt.Println(err)
	}
	buf := readSize(file, *read_amount)
	fmt.Println(string(buf))
}
