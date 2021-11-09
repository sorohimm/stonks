package quotes_interfaces

import "net/http"

type IStockRepo interface {
	GetQuotes(*http.Request) (interface{}, error)
}
