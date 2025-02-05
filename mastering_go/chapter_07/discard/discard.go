package main

import (
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		log.Println("Enabled logging!")
	} else {
		log.SetOutput(os.Stderr)
		log.Println("Disabling Logging!")
		log.SetOutput(io.Discard)
		log.Println("Does not print because it goes to io discard")
	}
}
