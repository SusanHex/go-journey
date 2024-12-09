package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	LOGFILE := path.Join(os.TempDir(), "yummygojuice.log")
	fmt.Println("Log FILE path:")
	fmt.Println(LOGFILE)
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	LstdFlags := log.Ldate | log.Lshortfile
	iLog := log.New(f, "LNum", LstdFlags)
	iLog.Println("Hello Beautiful World")
	iLog.SetFlags(log.Lshortfile | log.LstdFlags)
	iLog.Println("Hello Again, Beautiful")
}
