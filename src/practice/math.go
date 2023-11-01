package main

import (
	"fmt"
	"math"
)

func main() {
	var a int64 = 30
	var b float64 = 13
	var result float64 = float64(a) + b
	result = math.Round(result)
	fmt.Println("Result is", result)

	var realNumber float64 = 13.456
	fmt.Printf("Formatted number is %.2f\n", realNumber)
}
