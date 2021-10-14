package earnings_interfaces

import (
	"net/http"
	"stonks/internal/models/earnings"
)

type IEarningsRepo interface {
	GetEarnings(*http.Request) (earnings_model.Earnings, error)
}
