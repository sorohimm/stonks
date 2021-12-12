package growth_repo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"net/http"
	gmodels "stonks/internal/models/growth"
)

type GrowthRepo struct {
	Log    *zap.SugaredLogger
	Client *http.Client
}

func (r *GrowthRepo) GetGrowth(db *mongo.Database, coll string, filter interface{}) (gmodels.Response, error) {
	var body gmodels.Response
	cursor, err := db.Collection(coll).Aggregate(context.TODO(), filter)
	if err != nil {
		r.Log.Infof("growth_repo :: GetGrowth :: aggregate :: %s", err)
		return gmodels.Response{}, err
	}

	for cursor.Next(context.TODO()) {
		if err = cursor.Decode(&body); err != nil {
			r.Log.Infof("growth_repo :: GetGrowth :: decode :: %s", err)
			return gmodels.Response{}, err
		}
		break
	}
	r.Log.Info(body)
	return body, nil
}
