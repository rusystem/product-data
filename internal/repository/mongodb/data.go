package mongodb

import (
	"context"
	"github.com/rusystem/product-data/internal/config"
	"github.com/rusystem/product-data/pkg/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type Data struct {
	cfg *config.Config
	mdb *mongo.Database
}

func NewData(cfg *config.Config, mdb *mongo.Database) *Data {
	return &Data{
		cfg: cfg,
		mdb: mdb,
	}
}

func (r *Data) Insert(ctx context.Context, data []domain.Data) error {
	var docs []interface{}
	for _, t := range data {
		docs = append(docs, t)
	}

	_, err := r.mdb.Collection(r.cfg.MDB.Collection).InsertMany(ctx, docs)

	return err
}

func (r *Data) List(ctx context.Context, params domain.Params) ([]domain.Data, error) {

}
