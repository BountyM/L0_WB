package postgres

import (
	"github.com/BountyM/L0_WB/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func UpdateOrder(order models.Order, db *sqlx.DB) {
	query := `UPDATE orders SET track_number=$2, entry=$3, locale=$3, internal_signature=$3, customer_id=$3, 
	delivery_service=$3, shardkey=$3, sm_id=$3, date_created=$3, oof_shard=$3 WHERE order_uid=$1`
	_, err := db.Exec(query, order.Order_uid, order.Track_number, order.Entry, order.Locale, order.Internal_signature,
		order.Customer_id, order.Delivery_service, order.Shardkey, order.Sm_id, order.Date_created, order.Oof_shard)
	if err != nil {
		logrus.Info("error UpdateOrder postgres")
	}
	SetDelivery(order.Delivery, db, order.Order_uid)
	SetPayment(order.Payment, db)
	SetItems(order.Items, db)
}
