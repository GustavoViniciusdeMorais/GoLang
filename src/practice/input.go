package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text:")
	input, _ := reader.ReadString('\n')
	fmt.Println("Yout entered: ", input)

}