package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	min := 0
	max := 100
	count := 100
	seed := 0
	if len(os.Args) == 1 {
		fmt.Println("No arguments provided, using defaults")
	} else {
		if len(os.Args) < 4 {
			fmt.Println("Please provide at least three arguments. Please provide <min> <max> <count> [<seed>]")
			os.Exit(2)
		} else if len(os.Args) > 5 {
			fmt.Println("Too many Arguments provided. please provide <min> <max> <count> [<seed>]")
			os.Exit(1)
		} else if len(os.Args) == 5 {
			temp_seed, error := strconv.Atoi(os.Args[4])
			if error != nil {
				fmt.Println("Error parsing as seed an integer")
				os.Exit(5)
			}
			seed = temp_seed
			rand.Seed(int64(seed))
		}
		temp_min, error := strconv.Atoi(os.Args[1])
		if error != nil {
			fmt.Println("Error parsing min as an integer")
			os.Exit(3)
		}
		temp_max, error := strconv.Atoi(os.Args[2])
		if error != nil {
			fmt.Println("Error parsing max as an integer")
			os.Exit(4)
		}
		temp_count, error := strconv.Atoi(os.Args[3])
		if error != nil {
			fmt.Println("Error parsing count as an integer")
			os.Exit(5)
		}

		min = temp_min
		max = temp_max
		count = temp_count
	}
	for i := 0; i < count; i++ {
		fmt.Printf("%d ", random(min, max))
	}
	fmt.Println()
}
