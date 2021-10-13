package overview_interfaces

import (
	"net/url"
	"stonks/internal/models/overview"
)

type IOverviewService interface {
	GetOverview(url.Values) (overview_models.Overview, error)
}
