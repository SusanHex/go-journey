package main

import "fmt"

type aStructure struct {
	field1 complex128
	field2 int
}

func processPointer(x *float64) {
	*x = *x * *x
}

func returnPointer(x float64) *float64 {
	temp := 2 * x
	return &temp
}

func bothPointers(x *float64) *float64 {
	temp := 2 * *x
	return &temp
}

func main() {
	var f float64 = 12.123
	fmt.Println("Memory address of f:", &f)
	//pointer to f
	fP := &f
	fmt.Println("Memory address of f, (accessed through fP pointer):", fP)
	fmt.Println("Value of f:", *fP)
	processPointer(fP)
	fmt.Printf("Value of f: %.2f\n", f)
	//value of f does not change here
	x := returnPointer(f)
	//value of f does not change here
	fmt.Printf("value of x: %.2f\n", *x)
	xx := bothPointers(fP)
	fmt.Printf("Value of xx: %.2f\n", *xx)
	//check for empty structure
	var k *aStructure
	//this is nil because currently k points to nowhere
	fmt.Println(k)
	//therefore you are allowed to do this:
	if k == nil {
		k = new(aStructure)
	}
	fmt.Printf("%+v\n", k)
	if k != nil {
		fmt.Println("k is not nil!")
		fmt.Println(k.field1)
		fmt.Println(k.field2)
	}

}
