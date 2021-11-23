package choose_service

import (
	"go.uber.org/zap"
	"net/http"
	"net/url"
	"stonks/internal/config"
	m_const "stonks/internal/constants/market"
	"stonks/internal/db/filter"
	"stonks/internal/interfaces/api_interfaces"
	"stonks/internal/interfaces/choose_interfaces"
	"stonks/internal/interfaces/db_interfaces"
)

type ChooseService struct {
	Log           *zap.SugaredLogger
	Config        *config.Config
	StocksApiRepo api_interfaces.IQuotesApiRepo
	ChooseRepo    choose_interfaces.IChooseRepo
	DbRepo        db_interfaces.IDbRepo
	DbHandler     db_interfaces.IDBHandler
}

func (s *ChooseService) BuildRequest(values url.Values) *http.Request {
	values.Set("function", "TIME_SERIES_DAILY")
	values.Set("outputsize", "full")
	values.Set("apikey", s.Config.MarketKey)

	request, _ := http.NewRequest(http.MethodGet, m_const.URL, nil)
	request.URL.Path = m_const.Path
	request.URL.RawQuery = values.Encode()

	return request
}

func getColl(values url.Values) string {
	switch values.Get("interval") {
	case "daily":
		return "DailySeries"
	case "weekly":
		return "WeeklySeries"
	case "monthly":
		return "MonthlySeries"
	default:
		return ""
	}
}

func (s *ChooseService) GetChoose(values url.Values) (interface{}, error) {
	database := s.DbHandler.AcquireDatabase(s.Config.DbName)

	var pipe = filter.Choose(values)
	var err error
	var response interface{}

	//TODO add forecast choose
	switch values.Get("function") {
	case "PRICE":
		response, err = s.ChooseRepo.ChooseByPrice(database, getColl(values), pipe)
	case "PE":
		response, err = s.ChooseRepo.ChooseByPE(database, pipe)
	}

	if err != nil {
		s.Log.Infof("choose_service :: GetChoose :: %s", err)
		return nil, err
	}

	return response, nil
}
