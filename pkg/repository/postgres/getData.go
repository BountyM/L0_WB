package postgres

import (
	"github.com/BountyM/L0_WB/models"
	"github.com/jmoiron/sqlx"
)

func GetUser(db *sqlx.DB, user_id string) (models.Order, error) {
	query := `SELECT * FROM users WHERE "ID" = $1`
	var user models.Order
	err := db.Get(&user, query, user_id)
	return user, err
}

func GetUsers(db *sqlx.DB) ([]models.Order, error) {
	query := "SELECT * FROM users"
	var users []models.Order
	err := db.Select(&users, query)
	return users, err
}
