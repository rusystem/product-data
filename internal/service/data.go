package service

import (
	"context"
	"github.com/rusystem/product-data/internal/client"
	"github.com/rusystem/product-data/internal/repository"
	"github.com/rusystem/product-data/pkg/domain"
	"golang.org/x/exp/slices"
)

type DataService struct {
	repo          repository.Data
	productClient *client.Client
}

func NewDataService(repo repository.Data, productClient *client.Client) *DataService {
	return &DataService{
		repo:          repo,
		productClient: productClient,
	}
}

func (s *DataService) Fetch(ctx context.Context, url string) error {
	clientData, err := s.productClient.GetData(ctx, url)
	if err != nil {
		return err
	}

	mdbData, err := s.repo.GetAll(ctx)

	for _, v := range clientData {
		idx := slices.IndexFunc(mdbData, func(d domain.Data) bool { return d.Name == v.Name })

		if idx == 1 {
			if err := s.repo.UpdateOne(ctx, v); err != nil {
				return err
			}
		} else {
			if err := s.repo.InsertOne(ctx, v); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *DataService) List(ctx context.Context, params domain.Params) ([]domain.Data, error) {
	return s.repo.List(ctx, params)
}
