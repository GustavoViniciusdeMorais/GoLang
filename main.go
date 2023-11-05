package main

import (
	"fmt"

	"github.com/GustavoViniciusdeMorais/customer"
)

func main() {
	customer := customer.Customer{"Gustavo", 27}
	fmt.Println(customer.GetAge())
}
