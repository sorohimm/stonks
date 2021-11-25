package filter

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/url"
	"stonks/internal/models"
)

// CurrentPrice gives the current price for the symbol
func CurrentPrice(symbol string) mongo.Pipeline {
	p := mongo.Pipeline{
		{{"$match", bson.M{"_meta.symbol": symbol}}},
		{{"$project", bson.M{
			"symbol": "$_meta.symbol",
			"price":  bson.M{"$first": "$series.close"},
		},
		}},
	}

	return p
}


// Exist creates a filter that checks if a document with this symbol exists
func Exist(symbol string) interface{} {
	return bson.M{"_meta.symbol": symbol}
}

func ExistDetails(symbol string) interface{} {
	return bson.M{"symbol": symbol}
}


// Choose generates a pipeline for a choose request
func Choose(values url.Values) interface{} {
	var t models.ChooseTag
	t.Set(values)

	switch values.Get("by") {
	case "price":
		return price(t)
	case "pe":
		return pe(t)
	default:
		return nil
	}
}

// price is the helper method for Choose
func price(t models.ChooseTag) mongo.Pipeline {
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

// pe is the helper method for Choose
func pe(t models.ChooseTag) mongo.Pipeline {
	if t.Has("min") && t.Has("max") {
		p := mongo.Pipeline{
			{{"$match", bson.M{"trailingpe": bson.M{"$gte": t.Get("min"), "$lte": t.Get("max")}}}},
			{{
				"$project", bson.M{
					"symbol": "$symbol",
					"pe":     "$trailingpe",
				},
			}},
		}

		return p
	}

	if t.Has("min") && !t.Has("max") {
		p := mongo.Pipeline{
			{{"$match", bson.M{"trailingpe": bson.M{"$gte": t.Get("min")}}}},
			{{
				"$project", bson.M{
					"symbol": "$symbol",
					"pe":     "$trailingpe",
				},
			}},
		}

		return p
	}

	if !t.Has("min") && t.Has("max") {
		p := mongo.Pipeline{
			{{"$match", bson.M{"trailingpe": bson.M{"$lte": t.Get("max")}}}},
			{{
				"$project", bson.M{
					"symbol": "$symbol",
					"pe":     "$trailingpe",
				},
			}},
		}

		return p
	}

	return mongo.Pipeline{}
}

// Quotes generates a pipeline for a quote request
func Quotes(t models.Timing) interface{} {
	if t.Has("from") || t.Has("to") || t.Has("date") {
		return date(t)
	} else {
		return symbol(t)
	}
}

// symbol is the helper method for Quotes
func symbol(t models.Timing) mongo.Pipeline {
	return mongo.Pipeline{{{"$match", bson.M{"_meta.symbol": t.Get("symbol")}}}}
}

// date is the helper method for Quotes
func date(t models.Timing) mongo.Pipeline {
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

	if t.Has("from") && t.Has("to") && !t.Has("date") {
		p := mongo.Pipeline{
			{{"$match", bson.M{"_meta.symbol": t.Get("symbol")}}},
			{{"$project", bson.M{
				"series": bson.M{
					"$filter": bson.M{
						"input": "$series",
						"as":    "a",
						"cond": bson.M{
							"$and": bson.A{
								bson.M{"$gte": bson.A{
									"$$a.date", t.Get("from"),
								}},
								bson.M{"$lte": bson.A{
									"$$a.date", t.Get("to"),
								}},
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

// Growth generates a pipeline for a growth request
func Growth(t models.Timing) mongo.Pipeline {
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
			{{"$project", bson.M{
				"_meta": "$_meta",
				"begin": bson.M{"$first": "$begin"},
				"end":   bson.M{"$first": "$end"},
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
			"end": bson.M{"$first": "$series"},
		},
		}},
		{{"$project", bson.M{
			"_meta": "$_meta",
			"begin": bson.M{"$first": "$begin"},
			"end":   "$end",
		},
		}},
	}

	return p
}

func matchBySymbol(t models.Timing) mongo.Pipeline {
	return mongo.Pipeline{{{"$match", bson.M{"symbol": t.Get("symbol")}}}}
}
// Details generates a pipeline for a details request
func Details(values url.Values, function string) mongo.Pipeline {
	var t models.Timing
	t.Set(values)

	if function == "OVERVIEW" {
		return matchBySymbol(t)
	}

	if !t.Has("from") && !t.Has("to") && !t.Has("date") {
		return matchBySymbol(t)
	}

	if t.Has("from") && !t.Has("to") && !t.Has("date") {
		p := mongo.Pipeline{
			{{"$match", bson.M{"symbol": t.Get("symbol")}}},
			{{"$project", bson.M{
				fmt.Sprintf("%s", t.Get("interval")) : bson.M{
					"$filter": bson.M{
						"input": fmt.Sprintf("$%s", t.Get("interval")),
						"as":    "a",
						"cond": bson.M{
							"$gte": bson.A{
								"$$a.fiscalDateEnding", t.Get("from"),
							},
						},
					},
				}, "symbol": "$symbol",
			},
			}},
		}
		return p
	}

	if !t.Has("from") && t.Has("to") && !t.Has("date") {
		p := mongo.Pipeline{
			{{"$match", bson.M{"symbol": t.Get("symbol")}}},
			{{"$project", bson.M{
				fmt.Sprintf("%s", t.Get("interval")): bson.M{
					"$filter": bson.M{
						"input": fmt.Sprintf("$%s", t.Get("interval")),
						"as":    "a",
						"cond": bson.M{
							"$lte": bson.A{
								"$$a.fiscalDateEnding", t.Get("from"),
							},
						},
					},
				}, "symbol": "$symbol",
			},
			}},
		}
		return p
	}

	if t.Has("from") && t.Has("to") && !t.Has("date") {
		p := mongo.Pipeline{
			{{"$match", bson.M{"symbol": t.Get("symbol")}}},
			{{"$project", bson.M{
				fmt.Sprintf("%s", t.Get("interval")): bson.M{
					"$filter": bson.M{
						"input": fmt.Sprintf("$%s", t.Get("interval")),
						"as":    "a",
						"cond": bson.M{
							"$and": bson.A{
								bson.M{"$gte": bson.A{
									"$$a.fiscalDateEnding", t.Get("from"),
								}},
								bson.M{"$lte": bson.A{
									"$$a.fiscalDateEnding", t.Get("to"),
								}},
							},
						},
					},
				}, "symbol": "$symbol",
			},
			}},
		}
		return p
	}

	if !t.Has("from") && !t.Has("to") && t.Has("date") {
		p := mongo.Pipeline{
			{{"$match", bson.M{"symbol": t.Get("symbol")}}},
			{{"$project", bson.M{
				fmt.Sprintf("%s", t.Get("interval")): bson.M{
					"$filter": bson.M{
						"input": fmt.Sprintf("$%s", t.Get("interval")),
						"as":    "a",
						"cond": bson.M{
							"$eq": bson.A{
								"$$a.fiscalDateEnding", t.Get("from"),
							},
						},
					},
				}, "symbol": "$symbol",
			},
			}},
		}
		return p
	}

	return mongo.Pipeline{}
}