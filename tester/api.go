package main

import (
	"example.com/backend"
)

func main() {
	var app backend.App
	app.Initialize()
	app.Run()
}
