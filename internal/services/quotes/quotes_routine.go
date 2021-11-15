package quotes_service

import (
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"net/url"
	"stonks/internal/db"
)

var (
	collections         map[string]string
	intradayCollections map[string]string
)

func init() {
	collections = make(map[string]string)
	collections["TIME_SERIES_DAILY"] = "DailySeries"
	collections["TIME_SERIES_WEEKLY"] = "WeeklySeries"
	collections["TIME_SERIES_MONTHLY"] = "MonthlySeries"

	intradayCollections = make(map[string]string)
	intradayCollections["1min"] = "Intraday1"
	intradayCollections["5min"] = "Intraday5"
	intradayCollections["15min"] = "Intraday15"
	intradayCollections["30min"] = "Intraday30"
	intradayCollections["60min"] = "Intraday60"
}

func getCll(values url.Values) string {
	if values.Get("function") != "TIME_SERIES_INTRADAY" {
		return collections[values.Get("function")]
	} else {
		return intradayCollections[values.Get("interval")]
	}
}

func isDateRequest(v url.Values) bool {
	return v.Has("from") || v.Has("to") || v.Has("date")
}

func (s *QuotesService) DbQuotesRoutine(values url.Values, database *mongo.Database) (interface{}, error) {
	var err error
	filter, err := db.MakeFilter(values)
	if err != nil {
		s.Log.Errorf("make filter error: %s", err)
		return nil, err
	}

	var res interface{}

	switch values.Get("function") {
	case "TIME_SERIES_INTRADAY":
		res, err = s.QuotesRepo.GetIntraday(intradayCollections[values.Get("interval")], database, filter)
	case "TIME_SERIES_DAILY":
		if isDateRequest(values) {
			res, err = s.QuotesRepo.GetDaily("d", database, filter)
		} else {
			res, err = s.QuotesRepo.GetDaily("w", database, filter)
		}
	case "TIME_SERIES_WEEKLY":
		if isDateRequest(values) {
			res, err = s.QuotesRepo.GetWeekly("d", database, filter)
		} else {
			res, err = s.QuotesRepo.GetWeekly("w", database, filter)
		}
	case "TIME_SERIES_MONTHLY":
		if isDateRequest(values) {
			res, err = s.QuotesRepo.GetMonthly("d", database, filter)
		} else {
			res, err = s.QuotesRepo.GetMonthly("w", database, filter)
		}
	}

	if err != nil {
		s.Log.Infof("quotes_service: DbQuotesRoutine: %s", err)
		return nil, err
	}
	return res, nil
}

func (s *QuotesService) IntradayQuotesRoutine(r *http.Request) (interface{}, error) {
	var err error
	var res interface{}

	switch r.URL.Query().Get("interval") {
	case "1min":
		res, err = s.QuotesRepo.GetIntraday1Quotes(r)
	case "5min":
		res, err = s.QuotesRepo.GetIntraday5Quotes(r)
	case "15min":
		res, err = s.QuotesRepo.GetIntraday15Quotes(r)
	case "30min":
		res, err = s.QuotesRepo.GetIntraday30Quotes(r)
	case "60min":
		res, err = s.QuotesRepo.GetIntraday60Quotes(r)
	}
	if err != nil {
		s.Log.Infof("quotes_service: IntradayQuotesRoutine: %s", err)
		return nil, err
	}
	return res, nil
}

func (s *QuotesService) QuotesRoutine(values url.Values) (interface{}, error) {
	request := s.BuildRequest(values)
	var err error
	var res interface{}

	switch values.Get("function") {
	case "TIME_SERIES_INTRADAY":
		res, err = s.IntradayQuotesRoutine(request)
	case "TIME_SERIES_DAILY":
		res, err = s.QuotesRepo.GetDailyQuotes(request)
	case "TIME_SERIES_WEEKLY":
		res, err = s.QuotesRepo.GetWeeklyQuotes(request)
	case "TIME_SERIES_MONTHLY":
		res, err = s.QuotesRepo.GetMonthlyQuotes(request)
	}
	if err != nil {
		s.Log.Infof("quotes_service: QuotesRoutine: %s", err)
		return nil, err
	}
	return res, nil
}
