package quotes_interfaces

import (
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	qmodels "stonks/internal/models/quotes"
)

type IQuotesRepo interface {
	GetIntraday(string, *mongo.Database, interface{}) (qmodels.IntradayTSMongo, error)
	GetDaily(string, *mongo.Database, interface{}) (qmodels.DailyTSMongo, error)
	GetWeekly(string, *mongo.Database, interface{}) (qmodels.WeeklyTSMongo, error)
	GetMonthly(string, *mongo.Database, interface{}) (qmodels.MonthlyTSMongo, error)

	GetIntraday1Quotes(*http.Request) (qmodels.IntradayTSMongo, error)
	GetIntraday5Quotes(*http.Request) (qmodels.IntradayTSMongo, error)
	GetIntraday15Quotes(*http.Request) (qmodels.IntradayTSMongo, error)
	GetIntraday30Quotes(*http.Request) (qmodels.IntradayTSMongo, error)
	GetIntraday60Quotes(*http.Request) (qmodels.IntradayTSMongo, error)
	GetDailyQuotes(*http.Request) (qmodels.DailyTSMongo, error)
	GetWeeklyQuotes(*http.Request) (qmodels.WeeklyTSMongo, error)
	GetMonthlyQuotes(*http.Request) (qmodels.MonthlyTSMongo, error)

	InsertQuotes(string, *mongo.Database, interface{}) (interface{}, error)
}
