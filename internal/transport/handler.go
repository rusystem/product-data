package transport

import (
	"context"
	"github.com/rusystem/product-data/pkg/domain"
	data "github.com/rusystem/product-data/pkg/gen/data/proto"
)

func (h *Handler) Fetch(ctx context.Context, req *data.FetchRequest) (*data.Empty, error) {
	return &data.Empty{}, h.service.Data.Fetch(ctx, req.Url)
}

func (h *Handler) List(ctx context.Context, req *data.ListRequest) (*data.ListResponse, error) {
	param := domain.Params{
		Limit:  req.Limit,
		Entity: req.Entity.String(),
		Sort:   req.Sort.String(),
	}

	products, err := h.service.Data.List(ctx, param)
	if err != nil {
		return nil, err
	}

	var resp []*data.ListResponseEntity
	for _, product := range products {
		resp = append(resp, &data.ListResponseEntity{
			Name:      product.Name,
			Price:     product.Price,
			Changes:   product.Changes,
			Timestamp: product.Time,
		})
	}

	return &data.ListResponse{Products: resp}, nil
}
