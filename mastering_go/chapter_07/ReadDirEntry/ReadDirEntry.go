package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func GetSize(path string) (int64, error) {
	contents, err := os.ReadDir(path)
	if err != nil {
		return -1, err
	}
	var total int64
	for _, entry := range contents {
		if entry.IsDir() {
			temp, err := GetSize(filepath.Join(path, entry.Name()))
			if err != nil {
				return -1, err
			}
			total += temp
		} else {
			info, err := entry.Info()
			if err != nil {
				return -1, err
			}
			total += info.Size()
		}
	}
	return total, nil
}

func main() {

	if len(os.Args) == 1 {
		log.Fatal("Please provide a folder path")
	}
	length, err := GetSize(os.Args[1])
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf(`"%s" is %d bytes`+"\n", os.Args[1], length)
	}
}
