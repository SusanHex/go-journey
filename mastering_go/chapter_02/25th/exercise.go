package main

import "fmt"

func concatArraySlice(arr1, arr2 [10]int) []int {
	return append(arr1[:], arr2[:]...)
}
func concatArray(arr1, arr2 [10]int) [20]int {
	var tempArr [20]int
	for i := 0; i < 10; i++ {
		tempArr[i] = arr1[i]
		tempArr[i+10] = arr2[i]
	}
	return [20]int(tempArr)
}

func concatSliceArray(slice1, slice2 []int) [20]int {
	var tempArray [20]int
	for i := 0; i < 10; i++ {
		if i < len(slice1) {
			tempArray[i] = slice1[i]
		}
		if i < len(slice2) {
			tempArray[9+i] = slice2[i]
		}
	}
	return tempArray
}

func main() {
	arr1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	arr2 := [...]int{0, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println("arr1:", arr1)
	fmt.Println("arr2:", arr2)
	fmt.Println("concatArraySlice:", concatArraySlice(arr1, arr2))
	fmt.Println("concatArray:", concatArray(arr1, arr2))
	fmt.Println("concatSliceArray:", concatSliceArray(arr1[5:], arr2[:]))

}
