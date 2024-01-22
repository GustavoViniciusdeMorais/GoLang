package backend_test

import (
	"log"
	"os"
	"testing"

	"example.com/backend"
)

var backendApp backend.App

func TestMain(m *testing.M) {
	backendApp = backend.App{}
	backendApp.Initialize()
	CheckTablesExists()
	code := m.Run()
	os.Exit(code)
}

func CheckTablesExists() {
	_, err := backend.CreateTableProducts(backendApp.DB)
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	backendApp = backend.App{}
	backendApp.Initialize()
}
