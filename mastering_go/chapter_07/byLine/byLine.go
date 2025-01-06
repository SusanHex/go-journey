package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func lineByLine(input io.Reader) error {
	r := bufio.NewReader(input)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			break
		}
		if err != nil {
			fmt.Printf("error reading file: %s\n", err)
			return err
		}
		fmt.Print(line)
	}
	return nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide a file as input")
		os.Exit(1)
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	defer file.Close()
	lineByLine(file)
}
