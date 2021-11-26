package details_service

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"net/http"
	"net/url"
	"stonks/internal/config"
	"stonks/internal/constants/market"
	"stonks/internal/db"
	"stonks/internal/db/filter"
	"stonks/internal/interfaces/db_interfaces"
	"stonks/internal/interfaces/details_interfaces"
	"stonks/internal/interfaces/stocks_api_interfaces"
)

type CompanyDetailsService struct {
	Log           *zap.SugaredLogger
	Config        *config.Config
	StocksApiRepo stocks_api_interfaces.IStocksApiRepo
	DetailsRepo   details_interface.ICompanyDetailsRepo
	DbHandler     db_interfaces.IDBHandler
}

func (s *CompanyDetailsService) BuildRequest(values url.Values) *http.Request {
	values.Set("apikey", s.Config.MarketKey)

	request, _ := http.NewRequest(http.MethodGet, market_constants.URL, nil)
	request.URL.Path = market_constants.Path
	request.URL.RawQuery = values.Encode()

	return request
}

func (s *CompanyDetailsService) DbDetailsRoutine(database *mongo.Database, filter interface{}, function string) (interface{}, error) {
	var err error
	var res interface{}

	switch function {
	case "OVERVIEW":
		res, err = s.DetailsRepo.GetOverview(database, filter)
	case "EARNINGS":
		res, err = s.DetailsRepo.GetEarnings(database, filter)
	case "INCOME_STATEMENT":
		res, err = s.DetailsRepo.GetIncomeStatement(database, filter)
	case "BALANCE_SHEET":
		res, err = s.DetailsRepo.GetBalanceSheet(database, filter)
	case "CASH_FLOW":
		res, err = s.DetailsRepo.GetCashFlow(database, filter)
	}

	if err != nil {
		s.Log.Error("details_service :: DbDetailsRoutine :: cart error")
		return nil, err
	}

	return res, nil
}

func (s *CompanyDetailsService) GetCll(function string) string {
	//похуй на экосистему
	//сообщество какое-то
	//функциональное целое
	//связи между отдельными ... похуй
	switch function {
	case "OVERVIEW":
		return s.Config.DetailsCollections.Overview
	case "EARNINGS":
		return s.Config.DetailsCollections.Earnings
	case "INCOME_STATEMENT":
		return s.Config.DetailsCollections.IncomeStatement
	case "BALANCE_SHEET":
		return s.Config.DetailsCollections.BalanceSheet
	case "CASH_FLOW":
		return s.Config.DetailsCollections.CashFlow
	default:
		return ""
	}
}

func (s *CompanyDetailsService) GetCompanyDetails(values url.Values) (interface{}, error) {
	database := s.DbHandler.AcquireDatabase(s.Config.DbName)
	var err error
	var res interface{}
	f := filter.Details(values, values.Get("function"))
	s.Log.Info(f)
	if db.IsDocExist(database, s.GetCll(values.Get("function")), filter.ExistDetails(values.Get("symbol"))) {
		res, err =  s.DbDetailsRoutine(database, f, values.Get("function"))
		return res, nil
	}
		request := s.BuildRequest(values)
		switch values.Get("function") {
		case "OVERVIEW":
			res, err = s.StocksApiRepo.GetOverview(request)
		case "EARNINGS":
			res, err = s.StocksApiRepo.GetEarnings(request)
		case "INCOME_STATEMENT":
			res, err = s.StocksApiRepo.GetIncomeStatement(request)
		case "BALANCE_SHEET":
			res, err = s.StocksApiRepo.GetBalanceSheet(request)
		case "CASH_FLOW":
			res, err = s.StocksApiRepo.GetCashFlow(request)
		}
		if err != nil {
			s.Log.Error("details_service :: GetCompanyDetails :: cart error")
			return nil, err
		}

		_, err = s.DetailsRepo.InsertCompanyDetails(s.GetCll(values.Get("function")), database, res)

		res, err =  s.DbDetailsRoutine(database, f, values.Get("function"))
		if err != nil {
			s.Log.Error("details_service :: GetCompanyDetails :: cart error")
			return nil, err
		}

		return res, nil
}
