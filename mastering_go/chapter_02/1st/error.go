package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

var CustomError = fmt.Errorf("This is a custom Error message")

func check(a, b int) error {
	if a == 0 && b == 0 {
		return CustomError
	}
	return nil
}

func formattedError(a, b int) error {
	if a == 0 && b == 0 {
		return fmt.Errorf("a %d and b %d. UserID: %d", a, b, os.Getuid())
	}
	return nil
}

func main() {
	err := check(0, 10)
	if err == nil {
		fmt.Println("check() executed normally!")
	} else {
		fmt.Println(err)
	}
	err = check(0, 0)
	if errors.Is(err, CustomError) {
		fmt.Println("Custom error message detected")
	}
	err = formattedError(0, 0)
	if err != nil {
		fmt.Println(err)
	}
	i, err := strconv.Atoi("-456")
	if err == nil {
		fmt.Println("Int value is: ", i)
	}
	i, err = strconv.Atoi("Z456")
	if err != nil {
		fmt.Println(err)
	}
}
