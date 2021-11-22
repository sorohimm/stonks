package db_interfaces

import "go.mongodb.org/mongo-driver/mongo"

type IDbRepo interface {
	InsertOne(*mongo.Database, string, interface{}) (interface{}, error)
}
