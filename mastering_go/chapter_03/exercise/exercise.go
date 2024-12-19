package main

import "fmt"

func arrayToMap(arr [10]string) map[int]string {
	tempMap := make(map[int]string)
	for i, v := range arr {
		tempMap[i] = v
	}
	return tempMap
}

func mapToSlice(strMap map[string]string) ([]string, []string) {
	var keys, values []string
	for k, v := range strMap {
		keys = append(keys, k)
		values = append(values, v)
	}
	return keys, values
}

func main() {
	arr1 := [...]string{"Hello", "World", "Someone", "Let", "Me", "Out", "Of", "The", "Scary", "Matrix"}
	map1 := map[string]string{"Key1": "value1", "key2": "value2", "key3": "value3"}
	fmt.Println("arrayToMap:", arrayToMap(arr1))
	keys, values := mapToSlice(map1)
	fmt.Println("mapToSlice, Keys:", keys, "Values:", values)
}
