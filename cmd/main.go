package main

import (
	"context"
	"github.com/rusystem/product-data/internal/config"
	"github.com/rusystem/product-data/pkg/database"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
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

}
