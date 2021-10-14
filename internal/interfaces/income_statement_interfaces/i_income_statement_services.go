package income_statement_interfaces

import (
	"net/url"
	"stonks/internal/models/income_statement"
)


type IIncomeStatementService interface {
	GetIncomeStatement(url.Values) (income_statement_model.IncomeStatement, error)
}
