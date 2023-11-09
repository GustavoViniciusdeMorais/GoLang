package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	jsonString :=
		`[{"name":"apple","price":2.99,"quantity":3},
	{"name":"orange","price":1.50,"quantity":8},
	{"name":"banana","price":0.49,"quantity":12}]`

	result := getCartFromJson(jsonString)

	fmt.Println("GoLangGo!", result)
}

func getCartFromJson(jsonString string) []cartItem {
	var cart []cartItem

	decoder := json.NewDecoder(strings.NewReader(jsonString))
	_, err := decoder.Token()
	checkError(err)

	var catItem cartItem
	for decoder.More() {
		err := decoder.Decode(&catItem)
		checkError(err)
		cart = append(cart, catItem)
	}

	return cart
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type cartItem struct {
	Name     string
	Price    float64
	Quantity int
}
