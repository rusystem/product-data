package service

import "github.com/rusystem/product-data/internal/repository"

type DataService struct {
	repo repository.Data
}

func NewDataService(repo repository.Data) *DataService {
	return &DataService{
		repo: repo,
	}
}
