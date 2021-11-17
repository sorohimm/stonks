package growth_interfaces

import (
	"go.mongodb.org/mongo-driver/mongo"
	gmodels "stonks/internal/models/growth"
)

type IGrowthRepo interface {
	GetQuote(*mongo.Database, string, interface{}) (gmodels.Response, error)
}
