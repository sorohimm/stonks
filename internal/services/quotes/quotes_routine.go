package quotes_service

import (
	"net/http"
	"net/url"
	m_const "stonks/internal/constants/market"
	quotes_models "stonks/internal/models/quotes"
)

func (s *QuotesService) GetCll(values url.Values) string {
	if values.Get("function") == "TIME_SERIES_INTRADAY" {
		switch values.Get("interval") {
		case "1min":
			return s.Config.QuotesCollection.Intraday1
		case "5min":
			return s.Config.QuotesCollection.Intraday5
		case "15min":
			return s.Config.QuotesCollection.Intraday15
		case "30min":
			return s.Config.QuotesCollection.Intraday30
		case "60min":
			return s.Config.QuotesCollection.Intraday60
		default:
			return ""
		}
	} else {
		switch values.Get("function") {
		case "TIME_SERIES_DAILY":
			return s.Config.QuotesCollection.Daily
		case "TIME_SERIES_WEEKLY":
			return s.Config.QuotesCollection.Weekly
		case "TIME_SERIES_MONTHLY":
			return s.Config.QuotesCollection.Monthly
		default:
			return ""
		}
	}
}

func (s *QuotesService) FullRequest(values url.Values) *http.Request {
	if values.Get("function") == "TIME_SERIES_DAILY" {
		values.Set("outputsize", "full")
	}
	values.Set("apikey", s.Config.MarketKey)

	request, _ := http.NewRequest(http.MethodGet, m_const.URL, nil)
	request.URL.Path = m_const.Path
	request.URL.RawQuery = values.Encode()

	return request
}

func (s *QuotesService) IntradayQuotesRoutine(r *http.Request) (quotes_models.TSMongo, error) {
	var err error
	var res quotes_models.TSMongo

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
		return quotes_models.TSMongo{}, err
	}

	return res, nil
}

func (s *QuotesService) QuotesRoutine(request *http.Request) (quotes_models.TSMongo, error) {
	var err error
	var res quotes_models.TSMongo

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
		return quotes_models.TSMongo{}, err
	}

	return res, nil
}

func (s *QuotesService) SelectNewSeries(series quotes_models.TSMongo, from string) []quotes_models.SessionMongo {
	var res []quotes_models.SessionMongo
	for _, ticker := range series.Series {
		if ticker.Date >= from {
			res = append(res, ticker)
		}
	}

	return res
}
