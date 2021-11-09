package details_service

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
	"net/http"
	"net/url"
	"stonks/internal/config"
	"stonks/internal/constants/market"
	"stonks/internal/interfaces/db_interfaces"
	"stonks/internal/interfaces/details_interfaces"
)

type CompanyDetailsService struct {
	Log         *zap.SugaredLogger
	DetailsRepo details_interface.ICompanyDetailsRepo
	Config      *config.Config
	DbHandler   db_interfaces.IDBHandler
}

var collections map[string]string

func init() {
	collections = make(map[string]string)
	collections["OVERVIEW"] = "Overview"
	collections["EARNINGS"] = "Earnings"
	collections["INCOME_STATEMENT"] = "IncomeStatement"
	collections["BALANCE_SHEET"] = "BalanceSheet"
	collections["CASH_FLOW"] = "CashFlow"

}

func makeFilter(qp url.Values) bson.D {
	return bson.D{{"symbol", qp.Get("symbol")}}
}

func (s *CompanyDetailsService) GetCompanyDetails(qp url.Values) (interface{}, error) {
	db := s.DbHandler.AcquireDatabase(s.Config.DbName)

	collection := qp.Get("function")
	resp, err := s.DetailsRepo.GetDbCompanyDetails(qp.Get("function"), collection, db, makeFilter(qp))
	if err != nil {
		qp.Set("function", qp.Get("function"))
		qp.Set("apikey", s.Config.MarketKey)

		request, _ := http.NewRequest(http.MethodGet, market_constants.URL, nil)
		request.URL.Path = market_constants.Path
		request.URL.RawQuery = qp.Encode()

		resp, err := s.DetailsRepo.GetCompanyDetails(request)
		if err != nil {
			s.Log.Error("Vehicle error")
			return nil, err
		}

		_, err = s.DetailsRepo.InsertCompanyDetails(collection, db, resp)
		if err != nil {
			s.Log.Error("DB insert error")
			return nil, err
		}

		return resp, nil
	}

	return resp, nil
}
