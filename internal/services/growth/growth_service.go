package growth_services

import (
	"errors"
	"go.uber.org/zap"
	"net/http"
	"net/url"
	"stonks/internal/config"
	m_const "stonks/internal/constants/market"
	"stonks/internal/db"
	"stonks/internal/db/filter"
	"stonks/internal/interfaces/db_interfaces"
	"stonks/internal/interfaces/growth_interfaces"
	"stonks/internal/interfaces/stocks_api_interfaces"
	"stonks/internal/models"
	gmodels "stonks/internal/models/growth"
)

type GrowthService struct {
	Log           *zap.SugaredLogger
	Config        *config.Config
	GrowthRepo    growth_interfaces.IGrowthRepo
	QuotesApiRepo stocks_api_interfaces.IStocksApiRepo
	DbHandler     db_interfaces.IDBHandler
	DbRepo        db_interfaces.IDbRepo
}

func GrowthPercent(begin float64, end float64) float64 {
	return (end / begin) * 100 - 100
}

func CalculateGrowth(v *gmodels.Response) {
	v.Growth.OpenGrowth = GrowthPercent(v.Begin.Open, v.End.Open)
	v.Growth.HighGrowth = GrowthPercent(v.Begin.High, v.End.High)
	v.Growth.LowGrowth = GrowthPercent(v.Begin.Low, v.End.Low)
	v.Growth.CloseGrowth = GrowthPercent(v.Begin.Close, v.End.Close)
}

func (s *GrowthService) BuildRequest(values url.Values) *http.Request {
	values.Set("function", "TIME_SERIES_DAILY")
	values.Set("outputsize", "full")
	values.Set("apikey", s.Config.MarketKey)

	request, _ := http.NewRequest(http.MethodGet, m_const.URL, nil)
	request.URL.Path = m_const.Path
	request.URL.RawQuery = values.Encode()

	return request
}

func (s *GrowthService) GetGrowth(values url.Values) (interface{}, error) {
	database := s.DbHandler.AcquireDatabase(s.Config.DbName)
	var t models.Timing
	t.Set(values)

	var pipe = filter.Growth(t)

	if db.IsDocExist(database, s.Config.QuotesCollection.Daily, filter.Exist(t.Get("symbol"))) {
		response, err := s.GrowthRepo.GetGrowth(database, s.Config.QuotesCollection.Daily, pipe)
		if err != nil {
			s.Log.Info("growth_service :: GetGrowth :: dbroutine error")
			return nil, err
		}

		if response.Empty() {
			s.Log.Info("growth_service :: GetGrowth :: no suitable data")
			return nil, errors.New("there is no suitable data")
		}

		CalculateGrowth(&response)
		return response, nil
	} else {
		request := s.BuildRequest(values)
		apiResponse, err := s.QuotesApiRepo.GetDailyQuotes(request)
		if err != nil {
			s.Log.Infof("growth_service :: GetGrowth :: api error")
			return nil, err
		}

		_, err = s.DbRepo.InsertOne(database, s.Config.QuotesCollection.Daily, apiResponse)
		if err != nil {
			s.Log.Errorf("quotes repo :: GetGrowth :: InsertQuotes insertion error: %s", err)
			return nil, err
		}

		response, err := s.GrowthRepo.GetGrowth(database, s.Config.QuotesCollection.Daily, pipe)
		if err != nil {
			s.Log.Infof("growth_service :: GetGrowth :: dbroutine error")
			return nil, err
		}

		if response.Empty() {
			s.Log.Info("growth_service :: GetGrowth :: no suitable data")
			return nil, errors.New("there is no suitable data")
		}

		CalculateGrowth(&response)
		return response, nil
	}
}
