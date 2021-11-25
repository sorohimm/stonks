package db_interfaces

import (
	"go.mongodb.org/mongo-driver/mongo"
	"stonks/internal/models/choose"
)

type IDbRepo interface {
	InsertOne(*mongo.Database, string, interface{}) (interface{}, error)
	GetCurrentDailyPrice(database *mongo.Database, symbol string) (choose_models.Price, error)
}
