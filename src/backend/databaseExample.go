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
	db, err := sql.Open("sqlite3", "test.db")
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

func CreateSampleData(db *sql.DB) bool {
	statement := `
		INSERT INTO products (id , name, inventory, price ) 
		VALUES
		(1, 'test1', 1, 1.1),
		(2, 'test2', 2, 2.1);
	`

	_, err := db.Exec(statement)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Table products populated")

	return true
}

func ListProducts(db *sql.DB) bool {
	rows, err := db.Query("SELECT * FROM products")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id int64
		var name string
		var inventory int64
		var price float64
		err = rows.Scan(&id, &name, &inventory, &price)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d %s %d %f\n", id, name, inventory, price)
	}
	return true
}
