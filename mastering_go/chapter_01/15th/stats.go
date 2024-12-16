package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
)

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

func main() {
	arguments := os.Args
	data := make([]float64, len(arguments)-1)
	var min, max float64
	var initialized = 0
	nValues := 0
	var sum float64

	if len(arguments) == 1 {
		fmt.Println("No arguments provided, generating default")
		data = make([]float64, 10)
		nValues = 10
		for i := range data {
			data[i] = randomFloat(1, 100)
			sum += data[i]
			if initialized == 0 {
				min = data[i]
				max = data[i]
				initialized = 1
				continue
			}
			if data[i] < min {
				min = data[i]
			}
			if data[i] > max {
				max = data[i]
			}
		}

	}

	for i := 1; i < len(arguments); i++ {
		n, err := strconv.ParseFloat(arguments[i], 64)
		if err != nil {
			continue
		}
		nValues = nValues + 1
		sum = sum + n
		if initialized == 0 {
			min = n
			max = n
			initialized = 1
			continue
		}
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
		data[i-1] = n
	}
	fmt.Println("Number of values:", nValues)
	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
	if nValues == 0 {
		return
	}
	meanValue := sum / float64(nValues)
	fmt.Printf("Mean Value: %.5f\n", meanValue)
	// standard deviation
	var squared float64
	for i := 1; i < len(data); i++ {
		n := data[i]
		nValues = nValues + 1
		squared = squared + math.Pow((n-meanValue), 2)
	}
	standardDeviation := math.Sqrt(squared / float64(nValues))
	normalizedData := normalize(data, meanValue, standardDeviation)
	fmt.Printf("Standard Deviation: %.5f\n", standardDeviation)
	fmt.Printf("Normalized: %v\n", normalizedData)
}
