package main

import (
	"fmt"
	"os"
	"strconv"
)

type ar2x2 [2][2]int

func (a *ar2x2) Add(b ar2x2) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			a[i][j] = a[i][j] + b[i][j]
		}
	}
}

func (a *ar2x2) Subtract(b ar2x2) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			a[i][j] = a[i][j] - b[i][j]
		}
	}
}

func (a *ar2x2) Multiply(b ar2x2) {
	a[0][0] = a[0][0]*b[0][0] + a[0][1]*b[1][0]
	a[1][0] = a[1][0]*b[0][0] + a[1][1]*b[1][0]
	a[0][1] = a[0][0]*b[0][1] + a[0][1]*b[1][1]
	a[1][1] = a[1][0]*b[0][1] + a[1][1]*b[1][1]
}

func main() {
	if len(os.Args) != 9 {
		fmt.Println("Need 8 integers, sweetie!")
		os.Exit(1)
	}
	k := [8]int{}
	for index, i := range os.Args[1:] {
		v, err := strconv.Atoi(i)
		if err != nil {
			fmt.Printf("\"%s\" is not a valid integer, sweetie!", i)
			os.Exit(1)
		}
		k[index] = v
	}
	var a, b ar2x2
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			a[i][j] = k[i+j]
			b[i][j] = k[4+i+j]
		}
	}
	fmt.Println(a, b)
	fmt.Print(a, "*", b, "=")
	a.Multiply(b)
	fmt.Println(a)
	fmt.Print(a, "+", b, "=")
	a.Add(b)
	fmt.Println(a)
	fmt.Print(a, "-", b, "=")
	a.Subtract(b)
	fmt.Println(a)
}
