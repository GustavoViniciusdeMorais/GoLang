package main

import (
	"fmt"
)

func main() {
	anInt := 42
	var p = &anInt
	fmt.Println("memory address of p is ", p)
	fmt.Println("value of p is ", *p)

	*p = *p / 2
	fmt.Println("p value devided by 2 is ", *p)
}
