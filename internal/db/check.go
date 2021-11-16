package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

func IsDocExist(db *mongo.Database, coll string, filter interface{}) bool {
	err := db.Collection(coll).FindOne(context.TODO(), filter).Err()
	return err == nil
}
