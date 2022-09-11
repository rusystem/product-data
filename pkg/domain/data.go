package domain

import (
	"github.com/golang/protobuf/ptypes/timestamp"
)

type Data struct {
	Name    string               `bson:"name"`
	Price   int64                `bson:"price"`
	Changes int64                `bson:"changes"`
	Time    *timestamp.Timestamp `bson:"time"`
}
