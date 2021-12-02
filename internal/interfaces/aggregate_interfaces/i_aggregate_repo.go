package aggregate_interfaces

import (
	aggregate_models "stonks/internal/models/aggregate"
)

type IAggregateRepo interface {
	GetAggregate([]aggregate_models.Flow) float64
}
