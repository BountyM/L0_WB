package controllers

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/BountyM/L0_WB/models"
	"github.com/BountyM/L0_WB/pkg/repository/cache"
	"github.com/BountyM/L0_WB/pkg/repository/postgres"
	"github.com/jmoiron/sqlx"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
)

func MyMsgHandler(cache *cache.Cache, db *sqlx.DB) func(m *stan.Msg) {
	return func(m *stan.Msg) {
		jsonString := string(m.Data)
		fmt.Println(jsonString)
		var order models.Order
		err := json.Unmarshal([]byte(jsonString), &order)
		if err != nil {
			logrus.Fatalf("error json unmarshl :%s", err)
		} else {
			CheckCache(cache, order, db)
		}
	}
}

func CheckCache(cache *cache.Cache, order models.Order, db *sqlx.DB) {
	if _, err := cache.GetOrderByUid(order.Order_uid); err != nil {
		cache.Set(order.Order_uid, order)
		postgres.SetOrder(order, db)
	} else {
		order_db, err := postgres.GetOrder(db, order.Order_uid)
		if err != nil {
			logrus.Fatalf("error  postgres.GetUser in controllers CheckCache :%s", err)
		}
		if reflect.DeepEqual(order, order_db) {
			logrus.Info("duplicate key: " + order.Order_uid)
		} else {
			postgres.UpdateOrder(order, db)
		}
	}
}
