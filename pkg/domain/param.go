package domain

import "go.mongodb.org/mongo-driver/mongo/options"

type Params struct {
	Pagination Pagination
	Sort       Sort
}

type Pagination struct {
	Limit int64
}

type Sort struct {
	ByName            bool
	ByPriceAscending  bool
	ByPriceDescending bool
}

func (p Pagination) GetLimit() *int64 {
	if p.Limit == 0 {
		return nil
	}

	return &p.Limit
}

func GetParams(params *Params) *options.FindOptions {
	var opts *options.FindOptions
	if params != nil {
		opts = &options.FindOptions{
			Limit: params.Pagination.GetLimit(),
		}
	}

}
