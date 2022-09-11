package domain

import (
	"errors"
	data "github.com/rusystem/product-data/pkg/gen/data/proto"
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
