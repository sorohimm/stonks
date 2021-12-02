package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"stonks/internal/models"
)

func IsDocExist(db *mongo.Database, coll string, filter interface{}) bool {
	err := db.Collection(coll).FindOne(context.TODO(), filter).Err()
	return err == nil
}

func Relevance(db *mongo.Database, coll string, filter interface{}) (models.Relevance, error) {
	cursor, err := db.Collection(coll).Aggregate(context.TODO(), filter)
	if err != nil {
		return models.Relevance{}, err
	}

	body := models.Relevance{}
	for cursor.Next(context.TODO()) {
		if err = cursor.Decode(body); err != nil {
			return models.Relevance{}, nil
		}
	}

	return body, nil
}
