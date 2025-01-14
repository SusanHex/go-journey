package main

import (
	_ "embed"
	"fmt"
	"os"
)

//go:embed static/hello.txt
var f1 string

func writeToFile(s string, path string) error {
	fd, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer fd.Close()
	n, err := fd.Write([]byte(s))
	if err != nil {
		return err
	}
	fmt.Printf("Wrote %d bytes\n", n)
	return nil
}

func main() {
	fmt.Println("Length of embedded file", len(f1))
	fmt.Println(f1)
}
