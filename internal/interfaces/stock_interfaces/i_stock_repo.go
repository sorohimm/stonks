package stock_interfaces

import "net/http"

type IStockRepo interface {
	GetStock(*http.Request) (interface{}, error)
}