package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors = []string{"Red", "Green", "Blue"}
	colors = append(colors, "Black")
	fmt.Println("Arrays", colors)

	colors = append(colors[:len(colors)-1])
	fmt.Println("Removed last item", colors)

	numbers := make([]int, 5) // slice of five items
	numbers[0] = 221
	numbers[1] = 11
	sort.Ints(numbers) // sorted the slice of int values
	fmt.Println("Ordered", numbers)
}
