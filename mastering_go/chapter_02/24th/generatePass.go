package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
)

func generateBytes(n int64) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func main() {
	length := 0
	if len(os.Args) >= 2 {
		temp_length, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Printf("\"%s\" is not a valid number. please provide a number\n", os.Args[1])
			os.Exit(1)
		}
		length = temp_length
	} else {
		length = 8
	}
	byte_array, err := generateBytes(int64(length))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	generated_pass := base64.URLEncoding.EncodeToString(byte_array)
	fmt.Println(generated_pass)
}
