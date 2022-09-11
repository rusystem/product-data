package repository

import (
	"context"
	"github.com/rusystem/product-data/internal/config"
	"github.com/rusystem/product-data/internal/repository/mongodb"
	"github.com/rusystem/product-data/pkg/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type Data interface {
	UpdateOne(ctx context.Context, data domain.Data) error
	List(ctx context.Context, params domain.Params) ([]domain.Data, error)
}

type Repository struct {
	Data Data
}

func New(cfg *config.Config, mdb *mongo.Database) *Repository {
	return &Repository{
		Data: mongodb.NewDataRepository(cfg, mdb),
	}
}
