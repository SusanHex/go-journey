package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	flag := os.O_APPEND | os.O_CREATE | os.O_WRONLY
	file, err := os.OpenFile("gojuice.log", flag, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer file.Close()
	w := io.MultiWriter(file, os.Stderr)
	logger := log.New(w, "MyApp ", log.Lshortfile|log.Ldate|log.Ltime)
	logger.Printf("Process ID: %d", os.Getpid())
}
