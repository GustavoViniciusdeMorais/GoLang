package main

import (
	"fmt"
)

const pi float64 = 3.14

func main() {
	var text string = "This is Go!"
	fmt.Println(text)
	fmt.Printf("The variable typ is %T\n", text)

	var age int64 = 27
	fmt.Printf("My age is %d\n", age)

	withoutVar := 89
	fmt.Printf("Her age is %d\n", withoutVar)

	fmt.Printf("PI value is %f\n", pi)
	
}