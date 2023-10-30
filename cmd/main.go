package main

import (
	"time"

	l0wb "github.com/BountyM/L0_WB"
	"github.com/BountyM/L0_WB/controllers"
	"github.com/BountyM/L0_WB/models"
	"github.com/BountyM/L0_WB/pkg/handler"
	"github.com/BountyM/L0_WB/pkg/repository"
	"github.com/BountyM/L0_WB/pkg/repository/cache"
	"github.com/BountyM/L0_WB/pkg/repository/postgres"
	"github.com/BountyM/L0_WB/pkg/service"
	"github.com/joho/godotenv"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
)

func main() {
	cache := cache.NewCache() //cache

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := postgres.NewPostgresDB() //db
	if err != nil {
		logrus.Fatalf("failed to initialize db :%s", err.Error())
	}

	var users []models.Order

	// get all users from DB
	users, err = postgres.GetUsers(db)
	if err != nil {
		logrus.Fatalf("error by getting users from postgres :%s", err)
	}
	// set users into cache
	for _, user := range users {
		cache.Set(user.ID, user)
	}

	// nats
	// clusterID, clientID := "stanClusterID ", "clientID"

	sc, err := stan.Connect("test-cluster", "test", stan.NatsURL("nats://172.17.0.1:4223"))
	if err != nil {
		logrus.Infof("ERRFOJF : %s", err)
	}

	// Simple Async Subscriber
	preTime, _ := time.ParseDuration("1m")
	sub, _ := sc.Subscribe("foo", controllers.MyMsgHandler(cache, db), stan.StartAtTimeDelta(preTime))

	// run server
	repos := repository.NewRepository(cache)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(l0wb.Server)

	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}

	// Unsubscribe
	sub.Unsubscribe()

	// Close connection
	sc.Close()
}
