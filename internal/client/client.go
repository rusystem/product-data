package client

import (
	"context"
	"errors"
	"github.com/rusystem/product-data/pkg/domain"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Client struct {
	HTTPClient *http.Client
}

func New(timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("timeout can`t be zero")
	}

	return &Client{
		HTTPClient: &http.Client{
			Timeout: timeout,
		},
	}, nil
}

func (c *Client) GetData(ctx context.Context, url string) ([]domain.Data, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			return
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)

	var data []domain.Data
	for _, v := range strings.Split(string(body), "\n") {
		s := strings.Split(v, ";")
		if len(s) != 2 {
			break
		}

		price, err := strconv.ParseInt(s[1], 10, 64)
		if err != nil {
			return nil, err
		}

		data = append(data, domain.Data{
			Name:  s[0],
			Price: price,
			Time:  timestamppb.Now(),
		})
	}

	return data, err
}
