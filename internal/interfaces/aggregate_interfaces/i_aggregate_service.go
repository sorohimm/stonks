package aggregate_interfaces

import (
	"stonks/internal/models/aggregate"
)

type IAggregateService interface {
	GetAggregate(aggregate_models.Request) (interface{}, error)
}
