package aggregate_services

import (
	"errors"
	"go.uber.org/zap"
	"net/http"
	"net/url"
	"stonks/internal/config"
	m_const "stonks/internal/constants/market"
	"stonks/internal/db"
	"stonks/internal/db/filter"
	"stonks/internal/interfaces/aggregate_interfaces"
	"stonks/internal/interfaces/api_interfaces"
	"stonks/internal/interfaces/db_interfaces"
	"stonks/internal/models/aggregate"
)

type AggregateService struct {
	Log           *zap.SugaredLogger
	Config        *config.Config
	AggregateRepo aggregate_interfaces.IAggregateRepo
	StockApiRepo  api_interfaces.IStockApiRepo
	DbRepo        db_interfaces.IDbRepo
	DbHandler     db_interfaces.IDBHandler
}

func (s *AggregateService) BuildRequest(symbol string) *http.Request {
	var values url.Values
	values.Set("function", "TIME_SERIES_DAILY")
	values.Set("symbol", symbol)
	values.Set("outputsize", "full")
	values.Set("apikey", s.Config.MarketKey)

	request, _ := http.NewRequest(http.MethodGet, m_const.URL, nil)
	request.URL.Path = m_const.Path
	request.URL.RawQuery = values.Encode()

	return request
}

func (s *AggregateService) GetAggregate(request aggregate_models.Request) (interface{}, error) {
	database := s.DbHandler.AcquireDatabase(s.Config.DbName)

	for _, ticker := range request.Tickers {
		if !db.IsDocExist(database, "DailySeries", filter.Exist(ticker.Symbol)) {
			request := s.BuildRequest(ticker.Symbol)
			response, err := s.StockApiRepo.GetDailyQuotes(request)
			if err != nil {
				s.Log.Infof("quotes_service :: routine error")
				return nil, err
			}

			_, err = s.DbRepo.InsertOne(database, "DailySeries", response)
			if err != nil {
				s.Log.Infof("quotes_service :: database insert error")
				return nil, errors.New("server error")
			}
		}
	}

	var aggregate float64
	for _, ticker := range request.Tickers {
		currentValue, _ := s.DbRepo.GetCurrentDailyPrice(database, ticker.Symbol)
		aggregate += currentValue.Price * ticker.Coefficient
	}

	response := aggregate_models.Response{
		Tickers:   request.Tickers,
		Aggregate: aggregate,
	}

	return response, nil
}
