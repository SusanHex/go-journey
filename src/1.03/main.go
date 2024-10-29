package main

import (
	"fmt"
	"time"
)

var (
	Debug       bool
	LogLevel    string    = "info"
	startUpTime time.Time = time.Now()
)

func main() {
	fmt.Println(Debug, LogLevel, startUpTime)
}
