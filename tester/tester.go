package main

import (
	"example.com/backend"
	"example.com/practice"
)

func main() {
	practice.Test()
	backend.Run("localhost:9003")
}
