package quotes_repo

import (
	"context"
	_ "github.com/json-iterator/go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"net/http"
	qmodels "stonks/internal/models/quotes"
)

type QuotesRepo struct {
	Log    *zap.SugaredLogger
	Client *http.Client
}



func (r *QuotesRepo) GetQuotes(db *mongo.Database, coll string, filter interface{}) (interface{}, error) {
	var body qmodels.TSMongo
	cursor, err := db.Collection(coll).Aggregate(context.TODO(), filter)
	if err != nil {
		r.Log.Infof("quotes repo :: GetGrowth :: %s", err)
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		if err = cursor.Decode(&body); err != nil {
			return nil, err
		}
		break
	}

	return body, nil
}

func (r *QuotesRepo) UpdateQuotes(db *mongo.Database, coll string, filter interface{}, update interface{}) error {
	_ ,err := db.Collection(coll).UpdateMany(context.TODO(), filter, update)
	if err != nil {
		r.Log.Infof("quotes repo :: UpdateQuotes :: %s", err)
		return err
	}

	return nil
}