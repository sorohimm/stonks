package quotes_service

import (
	"net/http"
	"net/url"
	"stonks/internal/config"
	"stonks/internal/constants/market"
	"stonks/internal/interfaces/quotes_interfaces"
)

type QuotesService struct {
	StockRepo quotes_interfaces.IStockRepo
	Config    *config.Config
}

func (s *QuotesService) GetQuotes(queryParams url.Values) (interface{}, error) {
	queryParams.Set("function", queryParams.Get("function"))
	queryParams.Set("apikey", s.Config.MarketKey)

	request, _ := http.NewRequest(http.MethodGet, market_constants.URL, nil)
	request.URL.Path = market_constants.Path
	request.URL.RawQuery = queryParams.Encode()

	resp, err := s.StockRepo.GetQuotes(request)
	if err != nil {
		return err, nil
	}

	return resp, nil
}
