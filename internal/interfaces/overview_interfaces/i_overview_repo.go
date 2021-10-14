package overview_interfaces

import (
	"net/http"
	om "stonks/internal/models/overview"
)

type IOverviewRepo interface {
	GetOverview(*http.Request) (om.Overview, error)
}
