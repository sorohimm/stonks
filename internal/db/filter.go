package db

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/url"
	"stonks/internal/models"
)

type D bson.D

func makeDateFilter(t models.Timing) (interface{}, error) {
	if t.HasFrom() && t.HasTo() && t.HasDate() {
		return nil, errors.New("bad timing")
	}

	var m bson.M
	if t.HasFrom() && t.HasTo() && !t.HasDate() {
		m = bson.M{"$project": bson.M{
			"series": bson.M{
				"$filter": bson.M{
					"input": "$series",
					"as":    "a",
					"cond": bson.M{
						"$and": bson.M{
							"$lte": bson.M{
								"$$a.date": t.Get("to")},
							"$gte": bson.M{
								"$$a.date": t.Get("from"),
							},
						},
					},
				},
			},
			"symbol": t.Get("symbol"),
		},
		}
		pipeline := make([]bson.M, 0)

		pipeline = append(pipeline, m)
		return pipeline, nil
	}

	if t.HasFrom() && !t.HasTo() && !t.HasDate() {
		d := bson.D{{"$project", bson.M{
			"series": bson.M{
				"$filter": bson.M{
					"input": "$series",
					"as":    "a",
					"cond": bson.M{
						"$gte": bson.A{
							"$$a.date", t.Get("from"),
						},
					},
				},
			}, "symbol": t.Get("symbol")},
		},
		}

		return mongo.Pipeline{d}, nil
	}

	if !t.HasFrom() && t.HasTo() && !t.HasDate() {
		m = bson.M{"$project": bson.M{
			"series": bson.M{
				"$filter": bson.M{
					"input": "$series",
					"as":    "a",
					"cond": bson.M{
						"$lte": bson.M{
							"$$a.date": t.Get("from"),
						},
					},
				},
			},
			"symbol": t.Get("symbol"),
		},
		}
		pipeline := make([]bson.M, 0)

		pipeline = append(pipeline, m)

		return pipeline, nil
	}

	if !t.HasFrom() && !t.HasTo() && t.HasDate() {
		m = bson.M{"$project": bson.M{
			"series": bson.M{
				"$filter": bson.M{
					"input": "$series",
					"as":    "a",
					"cond": bson.M{
						"$eq": bson.M{
							"$$a.date": t.Get("from"),
						},
					},
				},
			},
			"symbol": t.Get("symbol"),
		},
		}
		pipeline := make([]bson.M, 0)

		pipeline = append(pipeline, m)

		return pipeline, nil
	}

	return m, nil
}

func MakeFilter(values url.Values) (interface{}, error) {
	t := models.Timing{}
	t.Set(values)

	if t.HasFrom() || t.HasTo() || t.HasDate() {
		filter, err := makeDateFilter(t)
		if err != nil {
			return nil, err
		}
		return filter, nil
	} else {
		return bson.M{"symbol": values.Get("symbol")}, nil
	}
}
