package src

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Inventory int64   `json:"inventory"`
	Price     float64 `json:"price"`
}

func GetProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("SELECT id, name, inventory, price FROM products")

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	products := []Product{}

	for rows.Next() {
		var p Product

		err = rows.Scan(&p.Id, &p.Name, &p.Inventory, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func (p *Product) GetProduct(db *sql.DB) error {
	return db.QueryRow("SELECT id, name, inventory, price FROM products WHERE id = ?", p.Id).
		Scan(&p.Id, &p.Name, &p.Inventory, &p.Price)
}

func (p *Product) CreateProduct(db *sql.DB) error {
	statement := `
		INSERT INTO products (name, inventory, price )
		VALUES
		(?, ?, ?);
	`

	res, err := db.Exec(statement, p.Name, p.Inventory, p.Price)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	p.Id = int(id)
	return nil
}