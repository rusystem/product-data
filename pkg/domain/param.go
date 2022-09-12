package domain

import (
	"errors"
	data "github.com/rusystem/product-data/pkg/gen/data/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
)

const (
	ENTITY_NAME  = "Name"
	ENTITY_PRICE = "Price"

	SORT_ASCENDING  = "Ascending"
	SORT_DESCENDING = "Descending"
)

var (
	entities = map[string]data.Entity{
		ENTITY_NAME:  data.Entity_Name,
		ENTITY_PRICE: data.Entity_Price,
	}

	sorts = map[string]data.Sort{
		SORT_ASCENDING:  data.Sort_Ascending,
		SORT_DESCENDING: data.Sort_Descending,
	}
)

type Params struct {
	Limit  int64
	Entity string
	Sort   string
}

func (p Params) GetLimit() *int64 {
	if p.Limit == 0 {
		return nil
	}

	return &p.Limit
}

func GetFindParams(params *Params, opts *options.FindOptions) *options.FindOptions {
	value := strings.ToLower(params.Entity)
	if params.Sort == "Ascending" {
		opts.SetSort(bson.M{value: 1})
	}

	if params.Sort == "Descending" {
		opts.SetSort(bson.M{value: -1})
	}

	if params.Limit > 0 {
		opts.SetLimit(params.Limit)
	}

	return opts
}

func ToPbEntity(entity string) (data.Entity, error) {
	val, ex := entities[entity]
	if !ex {
		return 0, errors.New("invalid entity")
	}

	return val, nil
}

func ToPbSort(action string) (data.Sort, error) {
	val, ex := sorts[action]
	if !ex {
		return 0, errors.New("invalid sorts")
	}

	return val, nil
}
