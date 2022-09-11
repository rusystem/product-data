package domain

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
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

func GetFindParams(params *Params) *options.FindOptions {
	var opts *options.FindOptions
	if params != nil {
		opts = &options.FindOptions{
			Limit: params.GetLimit(),
		}
	}

	value := strings.ToLower(params.Entity)

	if params.Sort == "Ascending" {
		opts.SetSort(bson.M{value: 1})
	}

	if params.Sort == "Descending" {
		opts.SetSort(bson.M{value: -1})
	}

	return opts
}
