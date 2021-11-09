package db_interfaces

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type IDBHandler interface {
	AcquireClient() *mongo.Client
	AcquireDatabase(string) *mongo.Database
}

