package choose_repo

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"net/http"
	"stonks/internal/models/choose"
)

type ChooseRepo struct {
	Log    *zap.SugaredLogger
	Client *http.Client
}

func (r *ChooseRepo) ChooseByPrice(database *mongo.Database, coll string, filter interface{}) (interface{}, error) {
	//var body choose_models.ByPrice
	var things []choose_models.Price
	cursor, err := database.Collection(coll).Aggregate(context.TODO(), filter)
	if err != nil {
		r.Log.Infof("choose_repo :: ChooseByPrice :: %s", err)
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

func (r *ChooseRepo) ChooseByPE(database *mongo.Database, filter interface{}) (interface{}, error) {
	var things []choose_models.PE
	cursor, err := database.Collection("Overview").Aggregate(context.TODO(), filter)
	if err != nil {
		r.Log.Infof("choose_repo :: ChooseByPrice :: %s", err)
		return nil, errors.New("aggregate error")
	}

	for cursor.Next(context.TODO()) {
		var this choose_models.PE
		if err = cursor.Decode(&this); err != nil {
			r.Log.Infof("choose_repo :: ChooseByPrice :: %s", err)
			return nil, errors.New("decode error")
		}
		things = append(things, this)
	}

	return things, nil
}
