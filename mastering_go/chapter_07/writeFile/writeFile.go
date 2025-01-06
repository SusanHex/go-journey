package main

import (
	"log"
	"os"
)

func main() {
	buffer := []byte("Data to write\n")
	f1, err := os.CreateTemp("", "temp_*")
	if err != nil {
		log.Fatal(err)
	}
	defer f1.Close()
	log.Println("Created temporary file:", f1.Name())
	n, err := f1.Write(buffer)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Wrote", n, "bytes to", f1.Name())
	_, err = f1.Seek(0, 0)
	if err != nil {
		log.Fatal(err)
	}
	read_buffer := make([]byte, n)
	_, err = f1.Read(read_buffer)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Read \"", string(read_buffer), "\"", "from file")
}
