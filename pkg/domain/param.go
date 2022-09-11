package domain

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Params struct {
	Pagination Pagination
	Sort       Sort
}

type Pagination struct {
	Limit int64 `form:"limit"`
}

type Sort struct {
	ByName           bool `form:"name"`
	ByPriceAscending bool `form:"price"`
}

func (p Pagination) GetLimit() *int64 {
	if p.Limit == 0 {
		return nil
	}

	return &p.Limit
}

func GetFindParams(params *Params) *options.FindOptions {
	var opts *options.FindOptions
	if params != nil {
		opts = &options.FindOptions{
			Limit: params.Pagination.GetLimit(),
		}
	}

	if params.Sort.ByName {
		opts.SetSort(bson.M{"name": 1})
	}

	if params.Sort.ByPriceAscending {
		opts.SetSort(bson.M{"price": 1})
	}

	return opts
}
