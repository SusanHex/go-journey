package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please provide a time")
	}
	input_time, err := time.Parse("02 January 2006 15:04", os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(input_time)
	MST, _ := time.LoadLocation("America/Denver")
	EST, _ := time.LoadLocation("America/New_York")
	fmt.Printf("MST: %v\n", input_time.In(MST))
	fmt.Printf("EST: %v\v", input_time.In(EST))

}
