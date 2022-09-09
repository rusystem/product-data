package database

import (
	"context"
	"github.com/rusystem/product-data/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoClient(ctx context.Context, cfg *config.Config) (*mongo.Client, error) {
	opts := options.Client()
	opts.SetAuth(options.Credential{
		Username: cfg.MDB.Username,
		Password: cfg.MDB.Password,
	})
	opts.ApplyURI(cfg.MDB.URI)

	dbClient, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	if err := dbClient.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return dbClient, nil
}
