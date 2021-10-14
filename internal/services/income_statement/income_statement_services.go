package income_statement_service

import (
	"fmt"
	"net/http"
	"net/url"
	"stonks/internal/config"
	"stonks/internal/constants/market"
	"stonks/internal/interfaces/income_statement_interfaces"
	ism "stonks/internal/models/income_statement"
)

type IncomeStatementService struct {
	IncomeStatementRepo income_statement_interfaces.IIncomeStatementRepo
	Config       *config.Config
}

func (s *IncomeStatementService) GetIncomeStatement(queryParams url.Values) (ism.IncomeStatement, error) {
	queryParams.Set("function", market_constants.IncomeStatement)
	queryParams.Set("apikey", s.Config.MarketKey)
	req, _ := http.NewRequest(http.MethodGet, market_constants.URL, nil)
	req.URL.Path = market_constants.Path
	req.URL.RawQuery = queryParams.Encode()
	sss := req.URL.String()
	fmt.Print(sss)
	resp, err := s.IncomeStatementRepo.GetIncomeStatement(req)
	if err != nil {
		return ism.IncomeStatement{}, err
	}

	return resp, nil
}
