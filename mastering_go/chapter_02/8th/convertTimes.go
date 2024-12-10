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
	input_time, err := time.Parse("RFC822", os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(input_time)
	loc, _ := time.LoadLocation("America/Denver")
	fmt.Printf("Denver Time: %s\n", loc)
}
