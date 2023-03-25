package backend

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type order struct {
	ID           int         `json: "id"`
	customerName string      `json: "customerName"`
	total        int         `json: "total"`
	status       string      `json: "status"`
	Items        []orderItem `json: "items"`
}

type orderItem struct {
	OrderId   int `json: "order_id"`
	ProductId int `json: "product_id"`
	Quantity  int `json: "quantity"`
}

func getOrders(db *sql.DB) ([]order, error) {
	rows, err := db.Query("SELECT * FROM orders")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	orders := []order{}

	for rows.Next() {
		var o order

		if err := rows.Scan(&o.ID, &o.customerName, &o.total, &o.status); err != nil {
			return nil, err
		}

		err = o.getOrderItems(db)

		if err != nil {
			return nil, err
		}

		orders = append(orders, o)
	}

	return orders, err
}

func (o *order) getOrderItems(db *sql.DB) error {
	rows, err := db.Query("SELECT * FROM order_items WHERE order_id = ?", o.ID)

	if err != nil {
		return err
	}

	defer rows.Close()

	orderItems := []orderItem{}

	for rows.Next() {
		var oi orderItem

		if err := rows.Scan(&oi.OrderId, &oi.ProductId, &oi.Quantity); err != nil {
			return err
		}

		orderItems = append(orderItems, oi)
	}

	o.Items = orderItems

	return nil
}

func (o *order) fetchOrder(db *sql.DB) error {
	db.QueryRow("SELECT customerName, total, status FROM orders WHERE id = ?", o.ID).Scan(&o.customerName, &o.total, &o.status)

	err := o.getOrderItems(db)

	if err != nil {
		return err
	}

	return nil
}

func (o *order) createOrder(db *sql.DB) error {
	res, err := db.Exec("INSERT INTO orders(customerName, total, status) VALUES (?,?,?)", o.customerName, o.total, o.status)

	if err != nil {
		return err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return err
	}

	o.ID = int(id)
	return nil
}

func (oi *orderItem) createOrderItems(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO order_itesm(order_id, product_id) VALUES (?, ?, ?)", oi.OrderId, oi.ProductId, oi.Quantity)

	if err != nil {
		return err
	}

	return nil
}
