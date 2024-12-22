package main

import "fmt"

func f1[S interface{ ~[]E }, E interface{}](x S) int {
	return len(x)
}

func f2[S ~[]E, E interface{}](x S) int {
	return len(x)
}

func f3[S ~[]E, E any](x S) int {
	return len(x)
}

func main() {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s2 := []string{"one", "two", "three"}
	s3 := []float64{6.1, 5.1, 6.7}
	fmt.Println(f1(s1), f2(s1), f3(s1))
	fmt.Println(f1(s2), f2(s2), f3(s3))
	fmt.Println(f1(s3), f2(s3), f3(s3))
}
