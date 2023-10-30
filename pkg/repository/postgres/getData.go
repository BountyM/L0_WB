package postgres

import (
	"github.com/BountyM/L0_WB/models"
	"github.com/jmoiron/sqlx"
)

func GetOrder(db *sqlx.DB, uid string) (models.Order, error) {
	query := `SELECT * FROM orders WHERE "order_uid" = $1`
	var order models.Order
	err := db.Get(&order, query, uid)
	return order, err
}
