package quotes_interfaces

import (
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	qmodels "stonks/internal/models/quotes"
)

type IQuotesRepo interface {
	GetIntraday(string, *mongo.Database, interface{}) (qmodels.IntradayTSMongo, error)

	GetQuotesDB(*mongo.Database, string, interface{}) (interface{}, error)

	GetIntraday1Quotes(*http.Request) (qmodels.TSMongo, error)
	GetIntraday5Quotes(*http.Request) (qmodels.TSMongo, error)
	GetIntraday15Quotes(*http.Request) (qmodels.TSMongo, error)
	GetIntraday30Quotes(*http.Request) (qmodels.TSMongo, error)
	GetIntraday60Quotes(*http.Request) (qmodels.TSMongo, error)
	GetDailyQuotes(*http.Request) (qmodels.TSMongo, error)
	GetWeeklyQuotes(*http.Request) (qmodels.TSMongo, error)
	GetMonthlyQuotes(*http.Request) (qmodels.TSMongo, error)

	InsertQuotes(string, *mongo.Database, interface{}) (interface{}, error)
}
