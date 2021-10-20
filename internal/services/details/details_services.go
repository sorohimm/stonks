package details_service

import (
	"net/http"
	"net/url"
	"stonks/internal/config"
	"stonks/internal/constants/details"
	"stonks/internal/interfaces/details_interfaces"
)

type CompanyDetailsService struct {
	DetailsRepo details_interface.ICompanyDetailsRepo
	Config       *config.Config
}

func (s *CompanyDetailsService) GetCompanyDetails(queryParams url.Values) (interface{}, error) {
	queryParams.Set("function", queryParams.Get("function"))
	queryParams.Set("apikey", s.Config.MarketKey)

	request, _ := http.NewRequest(http.MethodGet, market_constants.URL, nil)
	request.URL.Path = market_constants.Path
	request.URL.RawQuery = queryParams.Encode()

	resp, err := s.DetailsRepo.GetCompanyDetails(request)
	if err != nil {
		return err, nil
	}

	return resp, nil
}
