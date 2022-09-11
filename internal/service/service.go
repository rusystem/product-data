package service

import (
	"context"
	"github.com/rusystem/product-data/internal/client"
	"github.com/rusystem/product-data/internal/repository"
	"github.com/rusystem/product-data/pkg/domain"
)

type Data interface {
	Fetch(ctx context.Context, url string) error
	List(ctx context.Context, params domain.Params) ([]domain.Data, error)
}

type Service struct {
	Data Data
}

func New(repo *repository.Repository, productClient *client.Client) *Service {
	return &Service{
		Data: NewDataService(repo.Data, productClient),
	}
}
