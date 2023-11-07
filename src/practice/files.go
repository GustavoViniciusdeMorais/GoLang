package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "Hello from Go!"
	var path string = "./file.txt"
	file, err := os.Create(path)
	checkError(err)
	length, err := io.WriteString(file, content)
	checkError(err)
	fmt.Printf("Wrote %v characters to file\n", length)
	defer file.Close()
	defer readFile(path)
}

func readFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	checkError(err)
	fmt.Println("Text from file is: ", string(data))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
