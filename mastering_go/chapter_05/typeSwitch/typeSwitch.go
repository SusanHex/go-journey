package main

import "fmt"

type Secret struct {
	SecretValue string
}

type Entry struct {
	F1 int
	F2 string
	F3 Secret
}

func teststruct(x interface{}) {
	switch T := x.(type) {
	case Entry:
		fmt.Println("X is an Entry struct")
	case Secret:
		fmt.Println("X is Secret struct")
	default:
		fmt.Printf("Not Supported type: %T\n", T)
	}
}

func Learn(x interface{}) {
	switch T := x.(type) {
	default:
		fmt.Printf("Data type: %t\n", T)
	}
}

func main() {
	A := Entry{100, "F2", Secret{"Password123"}}
	teststruct(A)
	teststruct(A.F3)
	teststruct("A string")
	Learn(22.23)
	Learn('€')
}
