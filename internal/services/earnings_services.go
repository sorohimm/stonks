package services

import (
	"fmt"
	"net/http"
	"net/url"
	"stonks/internal/config"
	"stonks/internal/constants/market"
	"stonks/internal/interfaces/earnings_interfaces"
	"stonks/internal/models/earnings"
)

type EarningsService struct {
	EarningsRepo earnings_interfaces.IEarningsRepo
	Config       *config.Config
}

func (s *EarningsService) GetEarnings(queryParams url.Values) (earnings_models.Earnings, error) {
	queryParams.Set("function", market_constants.Earnings)
	queryParams.Set("apikey", s.Config.MarketKey)
	req, _ := http.NewRequest(http.MethodGet, market_constants.URL, nil)
	req.URL.Path = market_constants.Path
	req.URL.RawQuery = queryParams.Encode()
	sss := req.URL.String()
	fmt.Print(sss)
	resp, err := s.EarningsRepo.GetEarnings(req)
	if err != nil {
		return earnings_models.Earnings{}, err
	}

	return resp, nil
}
