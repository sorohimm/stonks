package overview_interfaces

import (
	"net/http"
	"stonks/internal/models/overview"
)

type IOverviewRepo interface {
	GetOverview(*http.Request) (overview_models.Overview, error)
}
