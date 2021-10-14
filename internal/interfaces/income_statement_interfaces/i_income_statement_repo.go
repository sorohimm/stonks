package income_statement_interfaces

import (
	"net/http"
	"stonks/internal/models/income_statement"
)


type IIncomeStatementRepo interface {
	GetIncomeStatement(*http.Request) (income_statement_model.IncomeStatement, error)
}
