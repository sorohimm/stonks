package db_repo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"net/http"
	"stonks/internal/db/filter"
	"stonks/internal/models/choose"
)

type DbRepo struct {
	Log    *zap.SugaredLogger
	Client *http.Client
}

// InsertOne executes an insert command to insert a single document into the collection.
func (r *DbRepo) InsertOne(database *mongo.Database, coll string, v interface{}) (interface{}, error) {
	res, err := database.Collection(coll).InsertOne(context.TODO(), v)
	if err != nil {
		r.Log.Errorf("db_repo :: InsertQuotes :: %s", err)
		return nil, err
	}
	return res.InsertedID, nil
}

// GetCurrentDailyPrice returns the current symbol price
func (r *DbRepo) GetCurrentDailyPrice(database *mongo.Database, symbol string) (interface{}, error) {
	f := filter.CurrentPrice(symbol)
	res := choose_models.Price{}
	cursor, err := database.Collection("DailySeries").Aggregate(context.TODO(), f)
	if err != nil {
		r.Log.Errorf("db_repo :: GetCurrentDailyPrice :: %s", err)
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		if err = cursor.Decode(&res); err != nil {
			r.Log.Errorf("db_repo :: GetCurrentDailyPrice :: %s", err)
			return choose_models.Price{}, err
		}
		break
	}

	return res, nil
}
