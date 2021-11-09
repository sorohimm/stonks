package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func IsHaveNecessaryData(db *mongo.Database, collection string, filter bson.D) bool {
	res := db.Collection(collection).FindOne(context.TODO(), filter)
	return res.Err() == nil
}
