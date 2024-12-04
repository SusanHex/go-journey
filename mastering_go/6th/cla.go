package main
import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Dang you! I need an argument! Nooooo!!!!")
		return
	}
}