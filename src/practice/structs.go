package main

import (
	"fmt"
)

func main() {
	viki := Cat{"Viki", 17}
	fmt.Println("Cat: ", viki)
	fmt.Printf("Object cat is %v\n", viki)

	fmt.Printf("The %s cat is %d years old\n", viki.Name, viki.Age)
}

// Cat is a struct
type Cat struct {
	Name string
	Age  int64
}
