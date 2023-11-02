package main

import (
	"fmt"
)

func main() {
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"

	fmt.Println("Arrays", colors)

	var numbers = [5]int64{1, 2, 3, 4, 5}
	fmt.Println("Numbers", numbers)

	fmt.Printf("Qtd of collors is %d, qtd of numbers is %d\n", len(colors), len(numbers))
}
