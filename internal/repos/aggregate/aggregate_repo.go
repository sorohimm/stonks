package aggregate_repo

import (
	"go.uber.org/zap"
	"net/http"
	aggregate_models "stonks/internal/models/aggregate"
)

type AggregateRepo struct {
	Log    *zap.SugaredLogger
	Client *http.Client
}

func (r *AggregateRepo) GetAggregate(tickers []aggregate_models.Flow) float64 {
	var aggregate float64
	for _, ticker := range tickers {
		aggregate += ticker.Price * ticker.Coeff
	}
	return aggregate
}
