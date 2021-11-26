package quotes_service

import (
	"errors"
	"go.uber.org/zap"
	"net/url"
	"stonks/internal/config"
	"stonks/internal/db"
	"stonks/internal/db/filter"
	"stonks/internal/interfaces/db_interfaces"
	"stonks/internal/interfaces/quotes_interfaces"
	"stonks/internal/interfaces/stocks_api_interfaces"
	"stonks/internal/models"
)

type QuotesService struct {
	Log           *zap.SugaredLogger
	Config        *config.Config
	QuotesRepo    quotes_interfaces.IQuotesRepo
	QuotesApiRepo stocks_api_interfaces.IStocksApiRepo
	DbRepo        db_interfaces.IDbRepo
	DbHandler     db_interfaces.IDBHandler
}

func (s *QuotesService) GetQuotes(values url.Values) (interface{}, error) {
	database := s.DbHandler.AcquireDatabase(s.Config.DbName)
	var coll = s.GetCll(values)
	var t models.Timing
	t.Set(values)
	var pipe = filter.Quotes(t)

	if db.IsDocExist(database, coll, filter.Exist(values.Get("symbol"))) {
		rel, _ := db.Relevance(database, coll, filter.Relevance(values.Get("symbol")))
		if ok := s.CheckRelevance(rel.LastRefreshed); !ok {
			request := s.FullRequest(values)
			response, err := s.QuotesRoutine(request)
			if err != nil {
				s.Log.Infof("quotes_service :: GetQuotes :: routine error")
				return nil, err
			}

			series := s.SelectNewSeries(response, rel.LastRefreshed)
			update := filter.UpdQuote(response.MetaData.LastRefreshed, series)
			match := filter.Match(values.Get("symbol"))

			err = s.QuotesRepo.Update(database, coll, match, update)
			if err != nil {
				return nil, err
			}
		}

		result, err := s.QuotesRepo.GetQuotes(database, coll, pipe)
		if err != nil {
			s.Log.Infof("quotes_service :: GetQuotes :: dbroutine error")
			return nil, err
		}
		return result, nil
	} else {
		request := s.FullRequest(values)
		response, err := s.QuotesRoutine(request)
		if err != nil {
			s.Log.Infof("quotes_service :: GetQuotes :: routine error")
			return nil, err
		}

		_, err = s.DbRepo.InsertOne(database, coll, response)
		if err != nil {
			s.Log.Infof("quotes_service :: GetQuotes :: database insert error")
			return nil, errors.New("server error")
		}

		if !t.HasInterval() {
			return response, nil
		}

		result, err := s.QuotesRepo.GetQuotes(database, coll, pipe)
		if err != nil {
			s.Log.Infof("quotes_service :: GetQuotes :: dbroutine error")
			return nil, err
		}

		return result, nil
	}
}
