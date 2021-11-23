package quotes_service

import (
	"errors"
	"go.uber.org/zap"
	"net/url"
	"stonks/internal/config"
	"stonks/internal/db"
	"stonks/internal/db/filter"
	"stonks/internal/interfaces/api_interfaces"
	"stonks/internal/interfaces/db_interfaces"
	"stonks/internal/interfaces/quotes_interfaces"
	"stonks/internal/models"
)

type QuotesService struct {
	Log           *zap.SugaredLogger
	Config        *config.Config
	QuotesRepo    quotes_interfaces.IQuotesRepo
	QuotesApiRepo api_interfaces.IQuotesApiRepo
	DbRepo        db_interfaces.IDbRepo
	DbHandler     db_interfaces.IDBHandler
}

func (s *QuotesService) GetQuotes(values url.Values) (interface{}, error) {
	database := s.DbHandler.AcquireDatabase(s.Config.DbName)
	var coll = getCll(values)
	var t models.Timing
	t.Set(values)
	var pipe = filter.Quotes(t)

	if db.IsDocExist(database, coll, filter.Exist(values.Get("symbol"))) {
		result, err := s.QuotesRepo.GetQuotes(database, coll, pipe)
		if err != nil {
			s.Log.Infof("quotes_service :: dbroutine error")
			return nil, err
		}
		return result, nil
	} else {
		request := s.BuildRequest(values)
		response, err := s.QuotesRoutine(request)
		if err != nil {
			s.Log.Infof("quotes_service :: routine error")
			return nil, err
		}

		_, err = s.DbRepo.InsertOne(database, coll, response)
		if err != nil {
			s.Log.Infof("quotes_service :: database insert error")
			return nil, errors.New("server error")
		}

		if !t.HasInterval() {
			return response, nil
		}

		result, err := s.QuotesRepo.GetQuotes(database, coll, pipe)
		if err != nil {
			s.Log.Infof("quotes_service :: dbroutine error")
			return nil, err
		}

		return result, nil
	}
}
