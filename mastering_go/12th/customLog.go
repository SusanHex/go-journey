package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	LOGFILE := path.Join(os.TempDir(), "yumgojuice.log")
	fmt.Println("Log file is: ", LOGFILE)
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	iLog := log.New(f, "iLog ", log.LstdFlags)
	iLog.Println("Hello beautiful world!")
	iLog.Println("Thank you for running me!")
}
