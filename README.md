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