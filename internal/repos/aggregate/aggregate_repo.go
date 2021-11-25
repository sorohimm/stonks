package aggregate_repo

import (
	"go.uber.org/zap"
	"net/http"
)

type AggregateRepo struct {
	Log    *zap.SugaredLogger
	Client *http.Client
}

func (r *AggregateRepo) GetAggregate() (interface{}, error) {
	return nil, nil
}
