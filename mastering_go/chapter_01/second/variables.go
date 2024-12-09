package main

import (
	"fmt"
	"math"
)

var Global int = 4000
var AnotherGlobal = -1

func main() {
	var j int
	i := Global + AnotherGlobal
	fmt.Println("Initial j value: ", j)
	j = Global
	k := math.Abs(float64(AnotherGlobal))
	fmt.Printf("Global=%d, i=%d, j=%d, k=%.2f\n", Global, i, j, k)
}
