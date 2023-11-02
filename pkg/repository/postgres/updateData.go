package postgres

import (
	"github.com/BountyM/L0_WB/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func UpdateOrder(order models.Order, db *sqlx.DB) {
	query := `UPDATE orders SET track_number=$2, entry=$3, locale=$4, internal_signature=$5, customer_id=$6, 
	delivery_service=$7, shardkey=$8, sm_id=$9, date_created=$10, oof_shard=$11 WHERE order_uid=$1`
	_, err := db.Exec(query, order.Order_uid, order.Track_number, order.Entry, order.Locale, order.Internal_signature,
		order.Customer_id, order.Delivery_service, order.Shardkey, order.Sm_id, order.Date_created, order.Oof_shard)
	if err != nil {
		logrus.Info("error UpdateOrder postgres")
	}

	if _, err := GetDelivery(db, order.Order_uid); err != nil {
		UpdateDelivery(order.Delivery, db, order.Order_uid)
	} else {
		SetDelivery(order.Delivery, db, order.Order_uid)
	}

	if _, err := GetPayment(db, order.Order_uid); err != nil {
		UpdatePayment(order.Delivery, db, order.Order_uid)
	} else {
		SetPayment(order.Payment, db)
	}

	if _, err := GetItems(db, order.Track_number); err != nil {
		UpdateItems(order.Items, db)
	} else {
		SetItems(order.Items, db)
	}
}

func UpdateDelivery(delivery models.Delivery, db *sqlx.DB, uid string) {
	query := `UPDATE deliveries SET name=$2, phone=$3, zip=$4, city=$5, address=$6, 
	region=$7, email=$8 WHERE order_uid=$1`
	_, err := db.Exec(query, uid, delivery.Name, delivery.Phone, delivery.Zip, delivery.City, delivery.Address, delivery.Region, delivery.Email)
	if err != nil {
		logrus.Info("error UpdateDelivery postgres")
	}
}

func UpdatePayment(delivery models.Delivery, db *sqlx.DB, uid string) {
	query := `UPDATE deliveries SET name=$2, phone=$3, zip=$4, city=$5, address=$6, 
	region=$7, email=$8 WHERE order_uid=$1`
	_, err := db.Exec(query, uid, delivery.Name, delivery.Phone, delivery.Zip, delivery.City, delivery.Address, delivery.Region, delivery.Email)
	if err != nil {
		logrus.Info("error UpdateDelivery postgres")
	}
}
