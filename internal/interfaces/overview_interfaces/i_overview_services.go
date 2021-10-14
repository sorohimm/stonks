package overview_interfaces

import (
	"net/url"
	om "stonks/internal/models/overview"
)

type IOverviewService interface {
	GetOverview(url.Values) (om.Overview, error)
}
