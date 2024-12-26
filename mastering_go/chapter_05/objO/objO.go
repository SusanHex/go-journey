package main

import "fmt"

type intA interface {
	foo()
}
type intB interface {
	bar()
}
type intC interface {
	intA
	intB
}

type a struct {
	XX int
	YY int
}
type b struct {
	AA string
	XX int
}
type c struct {
	A a
	B b
}

func processA(s intA) {
	fmt.Printf("%t\n", s)
}

// Satisfying intA
func (varC c) foo() {
	fmt.Println("Foo Processing", varC)
}

// Satisfying intB
func (varC c) bar() {
	fmt.Println("Bar Processing:", varC)
}

// structure compose gets the fields of structure a
type compose struct {
	field1 int
	a
}

func (A a) A() {
	fmt.Println("Function A() for A")
}
func (B b) A() {
	fmt.Println("Function A() for B")
}

func main() {
	var iC c = c{a{121, 34}, b{"-15", -15}}
	iC.A.A()
	iC.B.A()
	// following code will not work
	// iComp := compose{field1: 123, a{456, 789}}
	// iComp := compose{field1: 123, XX}
	iComp := compose{123, a{459, 789}}
	fmt.Println(iComp.XX, iComp.YY, iComp.field1)
	iC.bar()
	processA(iC)
}
