package grpc

import (
	"context"
	"fmt"
	"github.com/rusystem/product-data/pkg/domain"
	data "github.com/rusystem/product-data/pkg/gen/data/proto"
	"google.golang.org/grpc"
)

type Client struct {
	conn       *grpc.ClientConn
	dataClient data.DataClient
}

func NewClient(host string, port int) (*Client, error) {
	var conn *grpc.ClientConn

	addr := fmt.Sprintf("%s:%d", host, port)

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &Client{
		conn:       conn,
		dataClient: data.NewDataClient(conn),
	}, nil
}

func (c *Client) CloseConnection() error {
	return c.conn.Close()
}

func (c *Client) Fetch(ctx context.Context, url string) error {
	_, err := c.dataClient.Fetch(ctx, &data.FetchRequest{Url: url})

	return err
}

func (c *Client) List(ctx context.Context, req domain.Params) ([]domain.Data, error) {
	entity, err := domain.ToPbEntity(req.Entity)
	if err != nil {
		return nil, err
	}

	sort, err := domain.ToPbSort(req.Sort)
	if err != nil {
		return nil, err
	}

	params := &data.ListRequest{
		Limit:  req.Limit,
		Sort:   sort,
		Entity: entity,
	}

	resp, err := c.dataClient.List(ctx, params)
	if err != nil {
		return nil, err
	}

	var products []domain.Data
	for _, v := range resp.Products {
		products = append(products, domain.Data{
			Name:    v.Name,
			Price:   v.Price,
			Changes: v.Changes,
			Time:    v.Timestamp,
		})
	}

	return products, nil
}
