package controllers

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/BountyM/L0_WB/pkg/repository/cache"
	"github.com/jmoiron/sqlx"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
)

func MyMsgHandler(cache *cache.Cache, db *sqlx.DB) func(m *stan.Msg) {
	return func(m *stan.Msg) {
		jsonString := string(m.Data)
		fmt.Println(jsonString)
		var user models.User
		err := json.Unmarshal([]byte(jsonString), &user)
		if err != nil {
			logrus.Fatalf("error json unmarshl in myMsgHandler :%s", err)
		} else {
			CheckCache(cache, user, db)
		}
	}
}

func CheckCache(cache *repository.Cache, user models.User, db *sqlx.DB) {
	if _, err := cache.GetUserById(user.ID); err != nil {
		cache.Set(user.ID, user)
		fmt.Printf("cache : %s", cache.Users)
		postgres.SetUser(user, db)
	} else {
		fmt.Printf("user :%s", user)
		user_db, err := postgres.GetUser(db, user.ID)
		if err != nil {
			logrus.Fatalf("error  postgres.GetUser in controllers CheckCache :%s", err)
		}
		if reflect.DeepEqual(user, user_db) {
			logrus.Info("duplicate key: " + user.ID)
		} else {
			postgres.UpdateUser(user, db)
		}
	}
}
