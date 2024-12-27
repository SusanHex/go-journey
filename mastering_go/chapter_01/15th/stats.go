package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"slices"
	"sort"
	"strconv"
)

type DataFile struct {
	Filename string
	Len      int
	Minimum  float64
	Maximum  float64
	Mean     float64
	stdDev   float64
}

type DFslice []DataFile

func (a DFslice) Len() int {
	return len(a)
}
func (a DFslice) Less(i, j int) bool {
	return a[i].Mean < a[j].Mean
}
func (a DFslice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func normalize(data []float64, mean float64, stdDev float64) []float64 {
	if stdDev == 0 {
		return data
	}
	normalized := make([]float64, len(data))
	for i, val := range data {
		normalized[i] = math.Floor((val-mean)/stdDev*10000) / 10000
	}
	return normalized
}

func randomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func stdDev(x []float64) (float64, float64) {
	sum := 0.0
	for _, val := range x {
		sum = sum + val
	}
	meanValue := sum / float64(len(x))
	// fmt.Printf("Mean value: %.5f\n", meanValue)
	var squarded float64
	for i := 0; i < len(x); i++ {
		squarded = squarded + math.Pow((x[i]-meanValue), 2)
	}
	standardDeviation := math.Sqrt(squarded / float64(len(x)))
	return meanValue, standardDeviation
}

func readFile(filepath string) ([]float64, error) {
	_, err := os.Stat(filepath)
	if err != nil {
		return nil, err
	}
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}
	values := make([]float64, 0)
	for _, line := range lines {
		temp, err := strconv.ParseFloat(line[0], 64)
		if err != nil {
			log.Println("Error reading:", line[0], err)
		}
		values = append(values, temp)
	}
	return values, nil
}

func main() {
	if len(os.Args) == 1 {
		log.Println("Need one or more file paths, sweetie!")
		os.Exit(1)
	}
	files := DFslice{}

	for i := 1; i < len(os.Args); i++ {
		file := os.Args[i]
		currentFile := DataFile{}
		currentFile.Filename = file
		values, err := readFile(file)
		if err != nil {
			fmt.Println("Error reading:", file, err)
			os.Exit(0)
		}
		currentFile.Len = len(values)
		currentFile.Minimum = slices.Min(values)
		currentFile.Maximum = slices.Max(values)
		currentFile.Mean, currentFile.stdDev = stdDev(values)
		files = append(files, currentFile)
	}
	sort.Sort(files)
	for _, val := range files {
		f := val.Filename
		fmt.Println(f, ":", val.Len, val.Mean, val.Maximum, val.Minimum)
	}
}
