package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
)

func printWords(input_string string) {
	if len(input_string) == 0 {
		return
	}
	re := regexp.MustCompile(`[^\s]+`)
	for _, word := range re.FindAllString(input_string, -1) {
		fmt.Println(word)
	}
}

func wordByWord(input io.Reader) error {
	r := bufio.NewReader(input)
	for {
		line, err := r.ReadString('\n')
		if errors.Is(err, io.EOF) {
			printWords(line)
			break
		} else if err != nil {
			fmt.Println(err)
			return err
		}
		printWords(line)
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
	wordByWord(file)
}
