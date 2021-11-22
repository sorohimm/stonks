package quotes_service

import (
	"net/http"
	"net/url"
	m_const "stonks/internal/constants/market"

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

func (s *QuotesService) IntradayQuotesRoutine(r *http.Request) (interface{}, error) {
	var err error
	var res interface{}

	switch r.URL.Query().Get("interval") {
	case "1min":
		res, err = s.QuotesApiRepo.GetIntraday1Quotes(r)
	case "5min":
		res, err = s.QuotesApiRepo.GetIntraday5Quotes(r)
	case "15min":
		res, err = s.QuotesApiRepo.GetIntraday15Quotes(r)
	case "30min":
		res, err = s.QuotesApiRepo.GetIntraday30Quotes(r)
	case "60min":
		res, err = s.QuotesApiRepo.GetIntraday60Quotes(r)
	}
	if err != nil {
		s.Log.Infof("quotes_service :: IntradayQuotesRoutine :: %s", err)
		return nil, err
	}
	return res, nil
}

func (s *QuotesService) BuildRequest(values url.Values) *http.Request {
	if values.Get("function") == "TIME_SERIES_DAILY" {
		values.Set("outputsize", "full")
	}
	values.Set("apikey", s.Config.MarketKey)

	request, _ := http.NewRequest(http.MethodGet, m_const.URL, nil)
	request.URL.Path = m_const.Path
	request.URL.RawQuery = values.Encode()

	return request
}

func (s *QuotesService) QuotesRoutine(request *http.Request) (interface{}, error) {
	var err error
	var res interface{}

	switch request.URL.Query().Get("function") {
	case "TIME_SERIES_INTRADAY":
		res, err = s.IntradayQuotesRoutine(request)
	case "TIME_SERIES_DAILY":
		res, err = s.QuotesApiRepo.GetDailyQuotes(request)
	case "TIME_SERIES_WEEKLY":
		res, err = s.QuotesApiRepo.GetWeeklyQuotes(request)
	case "TIME_SERIES_MONTHLY":
		res, err = s.QuotesApiRepo.GetMonthlyQuotes(request)
	}
	if err != nil {
		s.Log.Infof("quotes_service :: QuotesRoutine :: %s", err)
		return nil, err
	}
	return res, nil
}
