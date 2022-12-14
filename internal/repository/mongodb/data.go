package mongodb

import (
	"context"
	"github.com/rusystem/product-data/internal/config"
	"github.com/rusystem/product-data/pkg/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DataRepository struct {
	cfg *config.Config
	mdb *mongo.Database
}

func NewDataRepository(cfg *config.Config, mdb *mongo.Database) *DataRepository {
	return &DataRepository{
		cfg: cfg,
		mdb: mdb,
	}
}

func (r *DataRepository) InsertOne(ctx context.Context, data domain.Data) error {
	_, err := r.mdb.Collection(r.cfg.MDB.Collection).InsertOne(ctx, data)

	return err
}

func (r *DataRepository) UpdateOne(ctx context.Context, data domain.Data) error {
	_, err := r.mdb.Collection(r.cfg.MDB.Collection).UpdateOne(ctx, bson.M{"name": data.Name}, bson.M{
		"$set": bson.M{
			"price": data.Price,
			"time":  data.Time,
		},
		"$inc": bson.M{
			"changes": 1,
		},
	}, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}

	return err
}

func (r *DataRepository) List(ctx context.Context, params domain.Params) ([]domain.Data, error) {
	findOptions := options.Find()
	opts := domain.GetFindParams(&params, findOptions)

	cur, err := r.mdb.Collection(r.cfg.MDB.Collection).Find(ctx, bson.D{}, opts)
	if err != nil {
		return nil, err
	}

	var data []domain.Data
	if err := cur.All(ctx, &data); err != nil {
		return nil, err
	}

	return data, err
}

func (r *DataRepository) GetAll(ctx context.Context) ([]domain.Data, error) {
	cur, err := r.mdb.Collection(r.cfg.MDB.Collection).Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	var data []domain.Data
	if err := cur.All(ctx, &data); err != nil {
		return nil, err
	}

	return data, err
}
