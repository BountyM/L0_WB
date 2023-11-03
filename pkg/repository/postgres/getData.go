package postgres

import (
	"errors"

	"github.com/BountyM/L0_WB/models"
	"github.com/jmoiron/sqlx"
)

func GetOrder(db *sqlx.DB, uid string) (models.Order, error) {
	var order models.Order
	query := `SELECT * FROM orders WHERE "order_uid" = $1`
	err := db.Get(&order, query, uid)
	if err != nil {
		return order, errors.New("the result set is empty, Orders")
	}
	order.Delivery, err = GetDelivery(db, uid)
	if err != nil {
		return order, errors.New("the result set is empty, Delivery")
	}
	order.Payment, err = GetPayment(db, uid)
	if err != nil {
		return order, errors.New("the result set is empty, Payment")
	}
	order.Items, err = GetItems(db, order.Track_number)
	if err != nil {
		return order, errors.New("the result set is empty, Items")
	}
	return order, err
}

func GetOrders(db *sqlx.DB) ([]models.Order, error) {
	orders := make([]models.Order, 1)
	query := `SELECT * FROM orders`
	rows, err := db.Query(query)
	if err != nil {
		return orders, err
	}

	for rows.Next() {
		var order models.Order
		err := rows.Scan(&order.Order_uid, &order.Track_number, &order.Entry, &order.Locale, &order.Internal_signature, &order.Customer_id,
			&order.Delivery_service, &order.Shardkey, &order.Sm_id, &order.Date_created, &order.Oof_shard)
		if err != nil {
			return orders, err
		}
		order.Delivery, err = GetDelivery(db, order.Order_uid)
		if err != nil {
			return orders, err
		}
		order.Payment, err = GetPayment(db, order.Order_uid)
		if err != nil {
			return orders, err
		}
		order.Items, err = GetItems(db, order.Track_number)
		if err != nil {
			return orders, err
		}
		orders = append(orders, order)
	}
	return orders, err
}

func GetDelivery(db *sqlx.DB, uid string) (models.Delivery, error) {
	var delivery models.Delivery
	query := `SELECT * FROM deliveries WHERE "order_uid" = $1`
	err := db.Get(&delivery, query, uid)
	return delivery, err
}

func GetPayment(db *sqlx.DB, uid string) (models.Payment, error) {
	var payment models.Payment
	query := `SELECT * FROM payments WHERE "transaction" = $1`
	err := db.Get(&payment, query, uid)

	return payment, err
}

func GetItems(db *sqlx.DB, track_number string) ([]models.Item, error) {
	items := make([]models.Item, 0)
	query := `SELECT * FROM items WHERE "track_number" = $1`
	err := db.Select(&items, query, track_number)

	return items, err
}
