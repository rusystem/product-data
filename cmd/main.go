package main

import (
	"context"
	"github.com/rusystem/product-data/internal/client"
	"github.com/rusystem/product-data/internal/config"
	"github.com/rusystem/product-data/internal/repository"
	"github.com/rusystem/product-data/internal/server"
	"github.com/rusystem/product-data/internal/service"
	"github.com/rusystem/product-data/internal/transport"
	"github.com/rusystem/product-data/pkg/database"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.ErrorLevel)
}

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	dbClient, err := database.NewMongoClient(ctx, cfg)
	if err != nil {
		logrus.Fatal(err)
	}
	defer func(dbClient *mongo.Client, ctx context.Context) {
		if err := dbClient.Disconnect(ctx); err != nil {
			return
		}
	}(dbClient, ctx)
	db := dbClient.Database(cfg.MDB.Database)

	productClient, err := client.New(cfg.Client.Timeout)
	if err != nil {
		logrus.Fatal(err)
	}

	dataRepo := repository.New(cfg, db)
	dataService := service.New(dataRepo, productClient)
	handler := transport.New(dataService)

	srv := server.New(handler)
	go func() {
		if err := srv.Run(cfg.Server.Host, cfg.Server.Port); err != nil {
			logrus.Fatal(err)
		}
	}()

	logrus.Info("Notes-log started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Info("Notes-log stopped")

	srv.Stop()

}
