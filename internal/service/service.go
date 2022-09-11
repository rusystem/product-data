package service

import (
	"context"
	"github.com/rusystem/product-data/internal/repository"
	data "github.com/rusystem/product-data/pkg/gen/data/proto"
	"google.golang.org/protobuf"
)

type Data interface {
	Fetch(ctx context.Context, req *data.FetchRequest) (, error)
	List(ctx context.Context)
}

type Service struct {
	Data Data
}

func New(repo *repository.Repository) *Service {
	return &Service{
		Data: NewDataService(repo.Data),
	}
}
