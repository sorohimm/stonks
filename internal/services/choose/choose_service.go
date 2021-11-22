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

func (s *ChooseService) GetChoose(values url.Values) (interface{}, error) {
	database := s.DbHandler.AcquireDatabase(s.Config.DbName)

	var t models.PriceTag
	t.Set(values)

	var pipe = filter.SymbolsByPrice(t)

	//TODO add pe and forecast choose
	//TODO choose by high low etc. when by==price
	response, err := s.ChooseRepo.GetChooseByPrice(database, "DailySeries", pipe)
	if err != nil {
		s.Log.Infof("choose_service :: database error")
		return nil, err
	}
	s.Log.Info(response)
	return response, nil
}
