package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	// "net/http/httptest"
)

var hello_list = []string{
	"Hello world",
	"今日は",
}

func main() {
	index := rand.Intn(len(hello_list))
	fmt.Println(len(hello_list))
	fmt.Println(hello_list[len(hello_list)-1])

	fmt.Println("Hello, world.")
	msg, err := hello(index)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)
}

func hello(index int) (string, error) {
	if index < 0 || index > len(hello_list)-1 {
		return "", errors.New("Out of range: " + strconv.Itoa(index))
	}
	return hello_list[index], nil
}
