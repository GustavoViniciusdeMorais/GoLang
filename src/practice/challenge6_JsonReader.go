package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonString :=
		`[{"name":"apple","price":2.99,"quantity":3},
	{"name":"orange","price":1.50,"quantity":8},
	{"name":"banana","price":0.49,"quantity":12}]`

	result := getCartFromJson(jsonString)

	fmt.Println("GoLangGo!", result)
}

// professor solution
func getCartFromJson(jsonString string) []cartItem {
	var cart []cartItem
	err := json.Unmarshal([]byte(jsonString), &cart)
	checkError(err)
	return cart
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

/* At this solution the struct attributes are
* directely mapped to the json attributes
 */
type cartItem struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}
