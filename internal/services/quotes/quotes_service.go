package quotes_service

import (
	"go.uber.org/zap"
	"net/http"
	"net/url"
	"stonks/internal/config"
	"stonks/internal/constants/market"
	"stonks/internal/interfaces/db_interfaces"
	"stonks/internal/interfaces/quotes_interfaces"
)

type QuotesService struct {
	Log        *zap.SugaredLogger
	Config     *config.Config
	QuotesRepo quotes_interfaces.IQuotesRepo
	DbHandler  db_interfaces.IDBHandler
}

func (s *QuotesService) BuildRequest(values url.Values) *http.Request {
	if values.Get("function") == "TIME_SERIES_DAILY" {
		values.Set("outputsize", "full")
	}
	values.Set("apikey", s.Config.MarketKey)

	request, _ := http.NewRequest(http.MethodGet, market_constants.URL, nil)
	request.URL.Path = market_constants.Path
	request.URL.RawQuery = values.Encode()

	return request
}

func (s *QuotesService) GetQuotes(values url.Values) (interface{}, error) {
	db := s.DbHandler.AcquireDatabase(s.Config.DbName)

	result, err := s.DbQuotesRoutine(values, db)
	if err != nil {
		response, err := s.QuotesRoutine(values)
		if err != nil {
			s.Log.Infof("quotes_service: routine error")
			return nil, err
		}

		collection := getCll(values)
		_, err = s.QuotesRepo.InsertQuotes(collection, db, response)
		if err != nil {
			s.Log.Infof("db insert error")
			return response, nil
		}

		res, err := s.DbQuotesRoutine(values, db)
		if err != nil {
			s.Log.Infof("get error")
			return response, nil
		}

		return res, nil
	}

	return result, nil
}
