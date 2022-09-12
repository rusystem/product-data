FROM golang:1.19.1-buster

RUN go version
ENV GOPATH=/

COPY ./ ./

# install psql
RUN apt-get update

# build go app
RUN go mod download
RUN go build -o product-data ./cmd/main.go

CMD ["./product-data"]