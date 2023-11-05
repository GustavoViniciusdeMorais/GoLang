package main

import (
	"fmt"
)

func main() {
	colors := []string{"Red", "Green", "Blue"}

	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	for i := range colors {
		fmt.Println(colors[i])
	}

	for _, color := range colors {
		fmt.Println(color)
	}

	value := 1
	for value < 10 {
		fmt.Println(value)
		value++
	}

	fmt.Println("GoLangGo!")
}
