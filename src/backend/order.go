package backend

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Order struct {
	Id           int         `json:"id"`
	CustomerName string      `json:"customerName"`
	Total        float64     `json:"total"`
	Items        []OrderItem `json:"items"`
}

type OrderItem struct {
	OrderId   int    `json:"orderId"`
	ProductId string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

func getOrders(db *sql.DB) ([]Order, error) {
	rows, err := db.Query("SELECT id, customerName, total FROM orders")

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	orders := []Order{}

	for rows.Next() {
		var o Order

		err = rows.Scan(&o.Id, &o.CustomerName, &o.Total)
		if err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}

	return orders, nil
}
