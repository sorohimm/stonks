package choose_repo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"net/http"
	"stonks/internal/models/choose"
)

type ChooseRepo struct {
	Log    *zap.SugaredLogger
	Client *http.Client
}

func (r *ChooseRepo) GetChooseByPrice(database *mongo.Database, coll string, filter interface{}) (interface{}, error) {
	//var body choose_models.ByPrice
	var things []choose_models.Price
	cursor, err := database.Collection(coll).Aggregate(context.TODO(), filter)
	if err != nil {
		r.Log.Infof("choose_repo :: GetChooseByPrice :: %s", err)
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		var this choose_models.Price
		if err = cursor.Decode(&this); err != nil {
			return nil, err
		}
		things = append(things, this)
	}

	return things, nil
}

