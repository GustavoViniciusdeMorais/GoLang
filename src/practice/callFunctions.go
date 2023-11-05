package main

import (
	"fmt"
)

func main() {
	sum, qty := addAllValues(5, 8)
	fmt.Printf("Sum is %d and length is %d\n", sum, qty)
}

func sumValues(a int, b int) int {
	return a + b
}

func addAllValues(values ...int) (int, int) {
	total := 0
	for _, v := range values {
		total += v
	}
	return total, len(values)
}
