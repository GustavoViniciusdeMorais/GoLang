package backend

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	Id        int64
	Name      string
	Inventory int64
	Price     float64
}

type DB struct{}

func DbConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}

	var version string
	err = db.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("SQLite version ", version)
	return db, err
}

func CloseDbConnection(db *sql.DB) bool {
	error := db.Close()

	if error != nil {
		log.Fatal(error)
	}
	return true
}

func CreateTableProducts(db *sql.DB) bool {
	statement := `
		DROP TABLE IF EXISTS products;
		CREATE TABLE products (id INTEGER PRIMARY KEY, name TEXT, inventory INT, price REAL);
	`

	_, err := db.Exec(statement)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Table products created")

	return true
}
