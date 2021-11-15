package details_interface

import (
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	dmodels "stonks/internal/models/company_details"
)

type ICompanyDetailsRepo interface {
	GetCompanyDetails(*http.Request) (interface{}, error)
	GetOverview(*mongo.Database, interface{}) (interface{}, error)
	GetEarnings(*mongo.Database, interface{}) (dmodels.Earnings, error)
	GetIncomeStatement(*mongo.Database, interface{}) (dmodels.IncomeStatement, error)
	GetBalanceSheet(*mongo.Database, interface{}) (dmodels.BalanceSheet, error)
	GetCashFlow(*mongo.Database, interface{}) (dmodels.CashFlow, error)
	InsertCompanyDetails(string, *mongo.Database, interface{}) (interface{}, error)
}
