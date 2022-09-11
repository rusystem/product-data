package service

import (
	"context"
	"github.com/rusystem/product-data/internal/repository"
)

type Data interface {
	Fetch(ctx context.Context)
	List(ctx context.Context)
}

type Service struct {
	Data Data
}

func New(repo *repository.Repository) *Service {
	return &Service{
		Data: NewData(repo.Data),
	}
}
