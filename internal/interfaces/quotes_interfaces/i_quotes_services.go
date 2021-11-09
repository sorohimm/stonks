package quotes_interfaces

import "net/url"

type IStockService interface {
	GetQuotes(url.Values) (interface{}, error)
}
