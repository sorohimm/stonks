package quotes_interfaces

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type IQuotesRepo interface {
	GetQuotes(*mongo.Database, string, interface{}) (interface{}, error)
	UpdateQuotes(db *mongo.Database, coll string, filter interface{}, update interface{}) error
}
