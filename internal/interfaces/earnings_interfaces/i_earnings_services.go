package earnings_interfaces

import (
	"net/url"
	"stonks/internal/models/earnings"
)

type IEarningsService interface {
	GetEarnings(url.Values) (earnings_models.Earnings, error)
}