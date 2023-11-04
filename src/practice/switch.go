package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())
	weekDay := rand.Intn(7) + 1
	var result string

	switch weekDay {
	case 1:
		result = "sumday"
	case 2:
		result = "monday"
	default:
		result = "other day"
	}

	fmt.Println("GoLangGo!", result)
}
