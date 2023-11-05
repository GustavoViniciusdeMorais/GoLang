package main

import (
	"fmt"
)

func main() {
	gustavo := Person{"Gustavo", 27, 1.65, 76}
	fmt.Println("GoLangGo!", gustavo.GetAge())
}

type Person struct {
	Name   string
	Age    int
	Hight  float64
	Weight float64
}

func (p Person) GetAge() int {
	return p.Age
}
