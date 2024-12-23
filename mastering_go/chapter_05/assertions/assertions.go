package main

import "fmt"

func returnNumber() interface{} {
	return 12
}

func main() {
	anInt := returnNumber()
	Number, ok := anInt.(int)
	if ok {
		fmt.Println("Type Assertion successful: ", Number)
	} else {
		fmt.Println("Type Assertion Failed!")
	}
	Number++
	fmt.Println(Number)
	//The next statement would fail because there is no type assertion to get the value
	// anInt++
	// The next statement fails, but the failure is under
	// control because of the ok bool variable that tells
	// whether the type assertion is successful or not.
	value, ok := anInt.(int64)
	if ok {
		fmt.Println("Type assertion successful: ", value)
	} else {
		fmt.Println("Type assertion failed")
	}
	// the next statement is successful, but
	// dangerious because it does not make sure that the type assurtion is successful
	// it just happens to be successful
	i := anInt.(int)
	fmt.Println(fmt.Println("i:", i))
	// the following will PANIC because anInt is not bool
	_ = anInt.(bool)
}
