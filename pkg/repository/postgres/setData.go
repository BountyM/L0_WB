package postgres

import (
	"github.com/BountyM/L0_WB/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func SetOrder(order models.Order, db *sqlx.DB) {

	query := `INSERT INTO orders (order_uid, track_number, entry, locale, internal_signature, customer_id, 
		delivery_service, shardkey, sm_id, date_created, oof_shard) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, err := db.Exec(query, order.Order_uid, order.Track_number, order.Entry, order.Locale, order.Internal_signature,
		order.Customer_id, order.Delivery_service, order.Shardkey, order.Sm_id, order.Date_created, order.Oof_shard)
	if err != nil {
		logrus.Info("error SetOrder postgres")
	}
	SetDelivery(order.Delivery, db, order.Order_uid)
	SetPayment(order.Payment, db)
	SetItems(order.Items, db)
}

func SetDelivery(delivery models.Delivery, db *sqlx.DB, order_uid string) {
	query := `INSERT INTO deliveries (name, phone, zip, city, address, region, email, order_uid) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := db.Exec(query, delivery.Name, delivery.Phone, delivery.Zip, delivery.City, delivery.Address, delivery.Region, delivery.Email, order_uid)
	if err != nil {
		logrus.Info("error SetDelivery postgres")
	}
}

func SetPayment(payment models.Payment, db *sqlx.DB) {
	query := `INSERT INTO payments (transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	_, err := db.Exec(query, payment.Transaction, payment.Request_id, payment.Currency, payment.Provider, payment.Amount, payment.Payment_dt,
		payment.Bank, payment.Delivery_cost, payment.Goods_total, payment.Custom_fee)
	if err != nil {
		logrus.Info("error SetPayment postgres")
	}
}

func SetItems(items []models.Item, db *sqlx.DB) {
	query := `INSERT INTO items (chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	for _, item := range items {
		_, err := db.Exec(query, item.Chrt_id, item.Track_number, item.Price, item.Rid, item.Name, item.Sale,
			item.Size, item.Total_price, item.Nm_id, item.Brand, item.Status)
		if err != nil {
			logrus.Info("error SetPayment postgres")
		}
	}
}
