package stock_interfaces

import "net/url"

type IStockService interface {
	GetStock(url.Values) (interface{}, error)
}
