package main

import "fmt"

func change(s []string) {
	s[0] = "change_function"
}

func main() {
	a := [4]string{"Hello", "wonderfully", "Beautiful", "World"}
	fmt.Println("a:", a)
	var S0 = a[:1]
	fmt.Println(S0)
	S0[0] = "S0"
	var S12 = a[1:3]
	fmt.Println(S12)
	S12[0] = "S12_0"
	S12[1] = "S12_1"
	fmt.Println("a:", a)
	fmt.Println("Capacity of S0:", cap(S0), "Length of S0", len(S0))
	S0 = append(S0, "N1")
	S0 = append(S0, "N2")
	S0 = append(S0, "N3")
	a[0] = "-N1"
	S0 = append(S0, "N4")
	fmt.Println("Capacity of S0:", cap(S0), "Length of S0", len(S0))
	a[0] = "-N1-"
	a[1] = "-N2-"
	fmt.Println("S0:", S0)
	fmt.Println("a: ", a)
	fmt.Println("S12: ", S12)
}
