package aggregate_interfaces

type IAggregateRepo interface {
	GetAggregate() (interface{}, error)
}
