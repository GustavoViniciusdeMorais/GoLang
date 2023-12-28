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
		err = o.getOrderItems(db)
		if err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}

	return orders, nil
}

func (order *Order) getOrder(db *sql.DB) error {
	db.QueryRow("SELECT id, customerName, total FROM orders WHERE id = ?", order.Id).
		Scan(&order.Id, &order.CustomerName, &order.Total)
	err := order.getOrderItems(db)
	if err != nil {
		return err
	}

	return nil
}

func (order *Order) getOrderItems(db *sql.DB) error {
	rows, err := db.
		Query("SELECT orderId, productId, quantity FROM orderItems WHERE orderId = ?", order.Id)
	if err != nil {
		return err
	}
	orderItems := []OrderItem{}
	for rows.Next() {
		var item OrderItem
		err := rows.Scan(&item.OrderId, &item.ProductId, &item.Quantity)
		if err != nil {
			orderItems = append(orderItems, item)
		}
	}
	order.Items = orderItems

	return nil
}

func (order *Order) createOrder(db *sql.DB) error {
	statement := `
		INSERT INTO orders (customerName, total)
		VALUES
		(?, ?);
	`

	res, err := db.Exec(statement, order.CustomerName, order.Total)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	order.Id = int(id)
	return nil
}

func (orderItem *OrderItem) createOrderItem(db *sql.DB) error {
	statement := `
		INSERT INTO orderItems (orderId, productId, quantity)
		VALUES
		(?, ?, ?);
	`
	_, err := db.Exec(statement, orderItem.OrderId, orderItem.ProductId, orderItem.Quantity)
	if err != nil {
		return err
	}
	return nil
}
