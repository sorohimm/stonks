package filter

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"stonks/internal/models"
)

func dataPipeline(t models.Timing) interface{} {
	if t.Has("from") && t.Has("to") && t.Has("date") {
		p := mongo.Pipeline{
			{{"$match", bson.M{"_meta.symbol": t.Get("symbol")}}},
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

func symbolPipeline(t models.Timing) mongo.Pipeline {
	return mongo.Pipeline{{{"$match", bson.M{"_meta.symbol": t.Get("symbol")}}}}
}

// CurrentPrice gives the current price for the symbol
func CurrentPrice(symbol string) mongo.Pipeline {
	p := mongo.Pipeline{
		{{"$match", bson.M{"_meta.symbol": symbol}}},
		{{"$project", bson.M{
			"symbol": "$_meta.symbol",
			"price":   bson.M{"$first": "$series.close"},
			},
		}},
	}

	return p
}

// SymbolsByPrice Creates a filter that chooses symbols by price
func SymbolsByPrice(t models.PriceTag) mongo.Pipeline {
	if t.Has("min") && t.Has("max") {
		p := mongo.Pipeline{
			{{"$match", bson.M{fmt.Sprintf("series.0.%s", t.Get("point")): bson.M{"$gte": t.Get("min"), "$lte": t.Get("max")}}}},
			{{
				"$project", bson.M{
					"symbol": "$_meta.symbol",
					"price":  bson.M{"$first": fmt.Sprintf("$series.%s", t.Get("point"))},
				},
			}},
		}

		return p
	}

	if t.Has("min") && !t.Has("max") {
		p := mongo.Pipeline{
			{{"$match", bson.M{fmt.Sprintf("series.0.%s", t.Get("point")): bson.M{"$gte": t.Get("min")}}}},
			{{
				"$project", bson.M{
					"symbol": "$_meta.symbol",
					"price":  bson.M{"$first": fmt.Sprintf("$series.%s", t.Get("point"))},
				},
			}},
		}

		return p
	}

	if !t.Has("min") && t.Has("max") {
		p := mongo.Pipeline{
			{{"$match", bson.M{fmt.Sprintf("series.0.%s", t.Get("point")): bson.M{"$lte": t.Get("max")}}}},
			{{
				"$project", bson.M{
					"symbol": "$_meta.symbol",
					"price":  bson.M{"$first": fmt.Sprintf("$series.%s", t.Get("point"))},
				},
			}},
		}

		return p
	}

	return mongo.Pipeline{}
}

// Exist Creates a filter that checks if a document with this symbol exists
func Exist(symbol string) interface{} {
	return bson.M{"_meta.symbol": symbol}
}

// QuotesPipeline generates a pipeline for a quote request
func QuotesPipeline(t models.Timing) interface{} {
	if t.Has("from") || t.Has("to") || t.Has("date") {
		return dataPipeline(t)
	} else {
		return symbolPipeline(t)
	}
}
