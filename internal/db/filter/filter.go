package filter

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"stonks/internal/models"
)

func dataPipeline(t models.Timing) interface{} {
	if t.Has("from") && t.Has("to") && t.Has("date") {
		p := mongo.Pipeline{
			{{"$match", bson.M{"_meta.symbol": t.Get("symbol")}}},
			{{"$project", bson.M{
					"series": bson.M{
						"$filter": bson.M{
							"input": "$series",
							"as":    "a",
							"cond": bson.M{
								"$and": bson.A{
									bson.M{"$lte": bson.A{
										"$$a.date", t.Get("to"),
									},
									},
									bson.M{"$gte": bson.A{
										"$$a.date", t.Get("from"),
									},
									},
								},
							},
						},
					}, "_meta": "$_meta",
				},
			}},
		}
		return p
	}

	if t.Has("from") && !t.Has("to") && !t.Has("date") {
		p := mongo.Pipeline{
			{{"$match", bson.M{"_meta.symbol": t.Get("symbol")}}},
			{{"$project", bson.M{
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
					}, "_meta": "$_meta",
				},
			}},
		}
		return p
	}

	if !t.Has("from") && t.Has("to") && !t.Has("date") {
		p := mongo.Pipeline{
			{{"$match", bson.M{"_meta.symbol": t.Get("symbol")}}},
			{{"$project", bson.M{
					"series": bson.M{
						"$filter": bson.M{
							"input": "$series",
							"as":    "a",
							"cond": bson.M{
								"$lte": bson.A{
									"$$a.date", t.Get("to"),
								},
							},
						},
					}, "_meta": "$_meta",
				},
			}},
		}
		return p
	}

	if !t.Has("from") && !t.Has("to") && t.Has("date") {
		p := mongo.Pipeline{
			{{"$match", bson.M{"_meta.symbol": t.Get("symbol")}}},
			{{"$project", bson.M{
					"series": bson.M{
						"$filter": bson.M{
							"input": "$series",
							"as":    "a",
							"cond": bson.M{
								"$eq": bson.A{
									"$$a.date", t.Get("date"),
								},
							},
						},
					},
					"_meta": "$_meta",
				},
			}},
		}
		return p
	}
	return mongo.Pipeline{}
}

func GrowthPipeline(t models.Timing) interface{} {
	if t.Has("to") {
		p := mongo.Pipeline{
			{{"$match", bson.M{"_meta.symbol": t.Get("symbol")}}},
			{{"$project", bson.M{
				"_meta": "$_meta",
				"begin": bson.M{
					"$filter": bson.M{
						"input": "$series",
						"as":    "a",
						"cond": bson.M{
							"$eq": bson.A{
								"$$a.date", t.Get("from"),
							},
						},
					},
				},
				"end": bson.M{
					"$filter": bson.M{
						"input": "$series",
						"as":    "a",
						"cond": bson.M{
							"$eq": bson.A{
								"$$a.date", t.Get("to"),
							},
						},
					},
				},
			},
			}},
			{{ "$project", bson.M{
				"_meta": "$_meta",
				"begin": bson.M{"$first": "$begin"},
				"end": 	 bson.M{"$first": "$end"},
			},
			}},
		}
		return p
	}

	p := mongo.Pipeline{
		{{"$match", bson.M{"_meta.symbol": t.Get("symbol")}}},
		{{"$project", bson.M{
				"_meta": "$_meta",
				"begin": bson.M{
					"$filter": bson.M{
						"input": "$series",
						"as":    "a",
						"cond": bson.M{
							"$eq": bson.A{
								"$$a.date", t.Get("from"),
							},
						},
					},
				},
				"end":   bson.M{"$first": "$series"},
				},
		}},
		{{ "$project", bson.M{
				 "_meta": "$_meta",
				 "begin": bson.M{"$first": "$begin"},
				 "end": "$end",
			},
		}},
	}

	return p
}

func symbolPipeline(t models.Timing) interface{} {
	return mongo.Pipeline{{{"$match", bson.M{"_meta.symbol": t.Get("symbol")}}}}
}

// Exist Creates a filter that checks if there is a document with this field and value
func Exist(s string) interface{} {
	return bson.M{"_meta.symbol": s}
}

func QuotesPipeline(t models.Timing) interface{} {
	if t.Has("from") || t.Has("to") || t.Has("date") {
		return dataPipeline(t)
	} else {
		return symbolPipeline(t)
	}
}
