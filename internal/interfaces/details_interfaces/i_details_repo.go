package details_interface

import (
	"go.mongodb.org/mongo-driver/mongo"
	dmodels "stonks/internal/models/company_details"
)

type ICompanyDetailsRepo interface {
	GetOverview(*mongo.Database, interface{}) (dmodels.OverviewMongo, error)
	GetEarnings(*mongo.Database, interface{}) (dmodels.EarningsMongo, error)
	GetIncomeStatement(*mongo.Database, interface{}) (dmodels.IncomeStatementMongo, error)
	GetBalanceSheet(*mongo.Database, interface{}) (dmodels.BalanceSheetMongo, error)
	GetCashFlow(*mongo.Database, interface{}) (dmodels.CashFlowMongo, error)
	InsertCompanyDetails(string, *mongo.Database, interface{}) (interface{}, error)
}
