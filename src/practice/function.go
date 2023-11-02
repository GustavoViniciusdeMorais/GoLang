package main

import (
	"bufio"
	"fmt"
	"os"
)

func gus(name string) {
	fmt.Printf("Name is %s", name)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text:")
	input, err := reader.ReadString('\n')

	if err != nil {
		panic("error")
	}

	gus(input)
}
