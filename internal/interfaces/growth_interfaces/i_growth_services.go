package growth_interfaces

import "net/url"

type IGrowthService interface {
	GetGrowth(url.Values) (interface{}, error)
}
