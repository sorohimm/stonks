package details_service

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"net/http"
	"net/url"
	"stonks/internal/config"
	"stonks/internal/constants/market"
	"stonks/internal/interfaces/api_interfaces"
	"stonks/internal/interfaces/db_interfaces"
	"stonks/internal/interfaces/details_interfaces"
)

type CompanyDetailsService struct {
	Log           *zap.SugaredLogger
	Config        *config.Config
	StocksApiRepo api_interfaces.IQuotesApiRepo
	DetailsRepo   details_interface.ICompanyDetailsRepo
	DbHandler     db_interfaces.IDBHandler
}

var (
	collections map[string]string
)

func init() {
	collections = make(map[string]string)
	collections["OVERVIEW"] = "Overview"
	collections["EARNINGS"] = "Earnings"
	collections["INCOME_STATEMENT"] = "IncomeStatement"
	collections["BALANCE_SHEET"] = "BalanceSheet"
	collections["CASH_FLOW"] = "CashFlow"
}

func MakeFilter(values url.Values) (interface{}, error) {
	switch values.Get("function") {
	case "OVERVIEW":
		return bson.D{{"symbol", values.Get("symbol")}}, nil
	case "EARNINGS":
		return bson.D{{"symbol", values.Get("symbol")}}, nil
	case "INCOME_STATEMENT":
		return bson.D{{"symbol", values.Get("symbol")}}, nil
	case "BALANCE_SHEET":
		return bson.D{{"symbol", values.Get("symbol")}}, nil
	case "CASH_FLOW":
		return bson.D{{"symbol", values.Get("symbol")}}, nil
	default:
		return nil, errors.New("bad function")
	}
}

func (s *CompanyDetailsService) BuildRequest(values url.Values) *http.Request {
	values.Set("apikey", s.Config.MarketKey)

	request, _ := http.NewRequest(http.MethodGet, market_constants.URL, nil)
	request.URL.Path = market_constants.Path
	request.URL.RawQuery = values.Encode()

	return request
}

func (s *CompanyDetailsService) DbGetDetailsRoutine(values url.Values, db *mongo.Database) (interface{}, error) {
	filter, err := MakeFilter(values)
	if err != nil {
		s.Log.Errorf("Make filter error: %s", err)
		return nil, err
	}

	switch values.Get("function") {
	case "OVERVIEW":
		res, err := s.DetailsRepo.GetOverview(db, filter)
		if err != nil {
			s.Log.Error("overview: no data in db")
			return nil, err
		}
		return res, nil
	case "EARNINGS":
		result, err := s.DetailsRepo.GetEarnings(db, filter)
		if err != nil {
			s.Log.Error("earnings: no data in db")
			return nil, err
		}
		if values.Has("timing") {
			res, err := result.ByTiming(values)
			if err != nil {
				s.Log.Error("earnings: bad timing")
				return result, nil
			}
			return res, nil
		}
		return result, nil
	case "INCOME_STATEMENT":
		result, err := s.DetailsRepo.GetIncomeStatement(db, filter)
		if err != nil {
			s.Log.Error("income_statement: no data in db")
			return nil, err
		}
		if values.Has("timing") {
			res, err := result.ByTiming(values)
			if err != nil {
				s.Log.Error("income_statement: bad timing")
				return result, nil
			}
			return res, nil
		}
		return result, nil
	case "BALANCE_SHEET":
		result, err := s.DetailsRepo.GetBalanceSheet(db, filter)
		if err != nil {
			s.Log.Error("balance_sheet: no data in db")
			return nil, err
		}
		if values.Has("timing") {
			res, err := result.ByTiming(values)
			if err != nil {
				s.Log.Error("balance_sheet: bad timing")
				return result, nil
			}
			return res, nil
		}
		return result, nil
	case "CASH_FLOW":
		result, err := s.DetailsRepo.GetCashFlow(db, filter)
		if err != nil {
			s.Log.Error("cash_flow: no data in db")
			return nil, err
		}
		if values.Has("timing") {
			res, err := result.ByTiming(values)
			if err != nil {
				s.Log.Error("cash_flow: bad timing")
				return result, nil
			}
			return res, nil
		}
		return result, nil
	}
	return nil, errors.New("unresolved error")
}

func (s *CompanyDetailsService) GetCompanyDetails(values url.Values) (interface{}, error) {
	database := s.DbHandler.AcquireDatabase(s.Config.DbName)

	result, err := s.DbGetDetailsRoutine(values, database)
	if err != nil {
		request := s.BuildRequest(values)
		if values.Get("function") == "OVERVIEW" {
			response, err := s.StocksApiRepo.GetOverview(request)
			if err != nil {
				s.Log.Error("api get error")
				return nil, err
			}

			s.Log.Info(response)
			_, err = s.DetailsRepo.InsertCompanyDetails(collections[values.Get("function")], database, response)
			if err != nil {
				s.Log.Error("database insert error")
				return nil, err
			}

			res, err := s.DbGetDetailsRoutine(values, database)
			if err != nil {
				s.Log.Error("get error")
				return nil, err
			}


			return res, err
		}

		response, err := s.DetailsRepo.GetCompanyDetails(request)
		if err != nil {
			s.Log.Error("api get error")
			return nil, err
		}

		_, err = s.DetailsRepo.InsertCompanyDetails(collections[values.Get("function")], database, response)
		if err != nil {
			s.Log.Error("database insert error")
			return response, nil
		}

		res, err := s.DbGetDetailsRoutine(values, database)
		if err != nil {
			s.Log.Error("get error")
			return nil, err
		}

		s.Log.Info("obtained from the api")
		return res, nil
	}
	s.Log.Info("obtained from the database")
	return result, nil
}
