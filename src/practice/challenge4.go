package main

import (
	"fmt"
)

func CalculateTotal(cart []cartItem) float64 {
	var total float64
	for _, v := range cart {
		itemTotal := v.price * float64(v.quantity)
		total += itemTotal
		fmt.Printf("Item %s qty %d price %f equals %f\n", v.name, v.quantity, v.price, itemTotal)
	}

	return total
}

func main() {
	var cart []cartItem
	var apples = cartItem{"apple", 2.99, 3}
	var oranges = cartItem{"orange", 1.50, 8}
	var bananas = cartItem{"banana", .49, 12}
	cart = append(cart, apples, oranges, bananas)
	result := CalculateTotal(cart)
	fmt.Println("Total cart value", result)
}

type cartItem struct {
	name     string
	price    float64
	quantity int
}
