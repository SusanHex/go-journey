package main

import (
	"fmt"
	"os"
)

var VERSION string

func main() {
	if len(os.Args) == 2 {
		if os.Args[1] == "Version" {
			fmt.Println("Version:", VERSION)
		}
	}
}
