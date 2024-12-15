package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

const MIN int = 0
const MAX int = 94

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func getString(len int64) string {
	temp := ""
	startChar := "!"
	var i int64 = 1
	for {
		myRand := random(MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		if i == len {
			break
		}
		i++
	}
	return temp
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
	fmt.Println(getString(int64(length)))
}
