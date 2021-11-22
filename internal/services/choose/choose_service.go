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
	"stonks/internal/models"
)

type ChooseService struct {
	Log           *zap.SugaredLogger
	Config        *config.Config
	QuotesApiRepo api_interfaces.IQuotesApiRepo
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

	var t models.PriceTag
	t.Set(values)

	var pipe = filter.SymbolsByPrice(t)

	var err error
	var response interface{}
	//TODO add pe and forecast choose
	if values.Get("by") == "price" {
		response, err = s.ChooseRepo.ChooseByPrice(database, getColl(values), pipe)
	} else if values.Get("by") == "pe" {

	}
	if err != nil {
		s.Log.Infof("choose_service :: database error")
		return nil, err
	}

	return response, nil
}
