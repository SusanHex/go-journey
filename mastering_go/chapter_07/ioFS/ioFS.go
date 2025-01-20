package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
)

//go:embed static/*
var f embed.FS
var searchString string = "second.txt"

func list(f embed.FS) error {
	return fs.WalkDir(f, ".", walkFunction)
}

func walkFunction(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	fmt.Printf("Path=%q, isDir=%v\n", path, d.IsDir())
	return nil
}

func extract(f embed.FS, filepath string) ([]byte, error) {
	s, err := fs.ReadFile(f, filepath)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func walkSearch(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	if d.Name() == searchString {
		fileInfo, err := fs.Stat(f, path)
		if err != nil {
			return err
		}
		fmt.Println("Found", path, "with size", fileInfo.Size())
	}
	return nil
}

func main() {
	// print all files in the embeded file system
	err := list(f)
	if err != nil {
		log.Fatal(err)
	}
	result, err := extract(f, "static/first.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("contents of first.txt:", string(result))
	static_dir, err := f.ReadDir("static")
	if err != nil {
		log.Fatal(err)
	}
	for _, dir := range static_dir {
		err = walkSearch("static/second.txt", dir, err)
		if err != nil {
			log.Fatal(err)
		}
	}
}
