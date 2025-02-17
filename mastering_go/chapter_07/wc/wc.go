package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func Count(input io.ReadSeeker) (line_count, word_count, rune_count, byte_count int, err error) {
	r := bufio.NewReader(input)

	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return -1, -1, -1, -1, err
		}
		line_count++
		word_count += CountWords(line)
	}
	_, err = input.Seek(0, 0)
	if err != nil {
		return -1, -1, -1, -1, err
	}
	rune_count, byte_count, err = CountRunesAndBytes(input)
	if err != nil {
		return -1, -1, -1, -1, err
	}
	return line_count, word_count, rune_count, byte_count, nil
}

func CountRunesAndBytes(input io.Reader) (rune_count int, byte_count int, err error) {
	r := bufio.NewReader(input)
	for {
		rune_size := 0
		_, rune_size, err = r.ReadRune()
		if err == io.EOF {
			byte_count += rune_size
			break
		}
		if err != nil {
			return -1, -1, err
		}
		byte_count += rune_size
		rune_count += 1
	}
	return rune_count, byte_count, nil
}

func CountWords(input string) int {
	re := regexp.MustCompile(`[^\s]+`)
	return len(re.FindAllString(string(input), -1))
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
	lines, words, runes, bytes, err := Count(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(lines, words, runes, bytes)
}
