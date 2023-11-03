package main

import (
	"fmt"
)

func convertToMap(items []string) map[string]float64 {
	result := make(map[string]float64)
	// Your code goes here
	qtyElements := len(items)
	var aFloat = 100.0 / float64(qtyElements)

	for _, v := range items {
		result[v] = aFloat
	}

	return result
}

func main() {
	fmt.Println("GoLangGo!")
	slice := []string{"apple", "banana", "orange", "date"}
	result := convertToMap(slice)
	fmt.Println(result)
}
