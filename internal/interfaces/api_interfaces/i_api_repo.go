package api_interfaces

import (
	"net/http"
	details_models "stonks/internal/models/company_details"
	qmodels "stonks/internal/models/quotes"
)

type IQuotesApiRepo interface {
	GetIntraday1Quotes(*http.Request) (qmodels.TSMongo, error)
	GetIntraday5Quotes(*http.Request) (qmodels.TSMongo, error)
	GetIntraday15Quotes(*http.Request) (qmodels.TSMongo, error)
	GetIntraday30Quotes(*http.Request) (qmodels.TSMongo, error)
	GetIntraday60Quotes(*http.Request) (qmodels.TSMongo, error)

	GetDailyQuotes(*http.Request) (qmodels.TSMongo, error)
	GetWeeklyQuotes(*http.Request) (qmodels.TSMongo, error)
	GetMonthlyQuotes(*http.Request) (qmodels.TSMongo, error)

	GetOverview(*http.Request) (details_models.OverviewMongo, error)
	GetEarnings(*http.Request) (details_models.EarningsMongo, error)
	GetCashFlow(*http.Request) (details_models.CashFlowMongo, error)
	GetIncomeStatement(*http.Request) (details_models.IncomeStatementMongo, error)
	GetBalanceSheet(*http.Request) (details_models.BalanceSheetMongo, error)
}
