package main

import (
	"fmt"
)

func main() {
	fmt.Println("GoLangGo!")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
