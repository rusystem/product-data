package domain

import "time"

type Data struct {
	Name    string    `bson:"name"`
	Price   int       `bson:"price"`
	Changes int       `bson:"changes"`
	Time    time.Time `bson:"time"`
}
