package main

import (
	"fmt"
)

func main() {
	answer := 42
	var result string

	if answer < 0 {
		result = "less then zero"
	} else if answer == 0 {
		result = "equals zero"
	} else {
		result = "greater then zero"
	}

	fmt.Println("GoLangGo!", result)
}
