# gRPC Server for fetch & get product information. 
### Before using you need to install:
1. docker-compose
2. make
3. golang
### Installing:
```
git clone https://github.com/rusystem/product-data.git
```
```
make build && make run
```

This project contains following grpc methods:

- Fetch(URL) - requests an external CSV file with a list of products at an external address.
  The CSV file should look like PRODUCT NAME;PRICE. The last price of each product is stored in the database with the date of the request. The number of product price changes is also saved.
- List(paging params, sorting params) - gets a page-by-page list of products with their
  prices, the number of price changes and the dates of their last update.
  There is a possibility of sorting.

### Then create a new main.go file in any other category and run the following commands:
```
go mod init example
```
```
go get -u github.com/rusystem/product-data.git
```

### Then add need dependency. 
1. google.golang.org/protobuf
1. google.golang.org/grpc
1. github.com/golang/protobuf
### Example:

```go
package main

import (
	"context"
	"fmt"
	"github.com/rusystem/product-data/pkg/client/grpc"
	"github.com/rusystem/product-data/pkg/domain"
	"log"
)

const (
	host = "localhost"
	port = 9000
	url  = "http://164.92.251.245:8080/api/v1/products/"
)

func main() {
	ctx := context.Background()

	client, err := grpc.NewClient(host, port)
	if err != nil {
		log.Fatal(err)
	}

	defer func(client *grpc.Client) {
		if err := client.CloseConnection(); err != nil {
			log.Fatal(err)
		}
	}(client)

	if err := client.Fetch(ctx, url); err != nil {
		return
	}

	products, err := client.List(ctx, domain.Params{
		Limit:  10,
		Entity: domain.ENTITY_NAME,
		Sort:   domain.SORT_ASCENDING,
	})
	fmt.Println(len(products))

	for _, product := range products {
		fmt.Printf("Name: %s, price: %d, changeCount: %d, changeTime: %v",
			product.Name, product.Price, product.Changes, product.Time)
	}
}

```