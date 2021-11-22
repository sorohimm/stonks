package growth_services

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"net/url"
	"stonks/internal/config"
	m_const "stonks/internal/constants/market"
	"stonks/internal/db"
	"stonks/internal/db/filter"
	"stonks/internal/interfaces/api_interfaces"
	"stonks/internal/interfaces/db_interfaces"
	"stonks/internal/interfaces/growth_interfaces"
	"stonks/internal/models"
	gmodels "stonks/internal/models/growth"
	"strconv"
)

type GrowthService struct {
	Log           *zap.SugaredLogger
	Config        *config.Config
	GrowthRepo    growth_interfaces.IGrowthRepo
	QuotesApiRepo api_interfaces.IQuotesApiRepo
	DbHandler     db_interfaces.IDBHandler
	DbRepo        db_interfaces.IDbRepo
}

func growthPercent(begin string, end string) string {
	bf, _ := strconv.ParseFloat(begin, 32)
	ef, _ := strconv.ParseFloat(end, 32)
	percent := (ef/bf)*100 - 100
	return fmt.Sprintf("%.3f%s", percent, " %")
}

func CalculateGrowth(v *gmodels.Response) {
	v.Growth.OpenGrowth = growthPercent(v.Begin.Open, v.End.Open)
	v.Growth.HighGrowth = growthPercent(v.Begin.High, v.End.High)
	v.Growth.LowGrowth = growthPercent(v.Begin.Low, v.End.Low)
	v.Growth.CloseGrowth = growthPercent(v.Begin.Close, v.End.Close)
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

	var pipe = filter.GrowthPipeline(t)

	if db.IsDocExist(database, "DailySeries", filter.Exist(t.Get("symbol"))) {
		response, err := s.GrowthRepo.GetGrowth(database, "DailySeries", pipe)
		if err != nil {
			s.Log.Infof("quotes_service :: dbroutine error")
			return nil, err
		}
		CalculateGrowth(&response)
		return response, nil
	} else {
		request := s.BuildRequest(values)
		apiResponse, err := s.QuotesApiRepo.GetDailyQuotes(request)
		if err != nil {
			s.Log.Infof("growth_service :: dbroutine error")
			return nil, err
		}

		_, err = s.DbRepo.InsertOne(database, "DailySeries", apiResponse)
		if err != nil {
			s.Log.Errorf("quotes repo: InsertQuotes insertion error: %s", err)
			return nil, err
		}

		response, err := s.GrowthRepo.GetGrowth(database, "DailySeries", pipe)
		if err != nil {
			s.Log.Infof("growth_service :: dbroutine error")
			return nil, err
		}
		return response, nil
	}
}
