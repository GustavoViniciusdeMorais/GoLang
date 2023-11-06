package main

import (
	"fmt"
	"strconv"
)

func calculate(input1 string, input2 string, operation string) float64 {
	a, errorA := convertInputToValue(input1)
	b, errorB := convertInputToValue(input2)

	if errorA != nil || errorB != nil {
		panic("Error!")
	}

	var result float64

	switch operation {
	case "+":
		result = addValues(a, b)
	case "-":
		result = subtractValues(a, b)
	case "/":
		result = divideValues(a, b)
	case "*":
		result = multiplyValues(a, b)
	}

	return result
}

func convertInputToValue(input string) (float64, error) {
	return strconv.ParseFloat(input, 8)
}

func addValues(value1, value2 float64) float64 {
	return value1 + value2
}

func subtractValues(value1, value2 float64) float64 {
	return value1 - value2
}

func multiplyValues(value1, value2 float64) float64 {
	return value1 * value2
}

func divideValues(value1, value2 float64) float64 {
	return value1 / value2
}

func main() {
	value1 := "10"
	value2 := "5.5"
	operation := "+"
	result := calculate(value1, value2, operation)
	fmt.Println("GoLangGo!", result)
}
