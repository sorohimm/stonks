package api_repo

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	qmodels "stonks/internal/models/quotes"
)

type ApiRepo struct {
	Log    *zap.SugaredLogger
	Client *http.Client
}

func (a *ApiRepo) GetIntraday1Quotes(request *http.Request) (qmodels.TSMongo, error) {
	resp, err := a.Client.Do(request)
	if err != nil || resp.StatusCode != 200 {
		a.Log.Errorf("api_repo :: GetIntraday1Quotes :: %s", err)
		return qmodels.TSMongo{}, err
	}

	body := qmodels.Intraday1TS{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&body)
	if err != nil {
		a.Log.Errorf("api_repo :: GetIntraday1Quotes :: %s", err)
		return qmodels.TSMongo{}, err
	}

	res := qmodels.TSMongo{}
	res.Set1(body)
	return res, nil
}

func (a *ApiRepo) GetIntraday5Quotes(request *http.Request) (qmodels.TSMongo, error) {
	resp, err := a.Client.Do(request)
	if err != nil || resp.StatusCode != 200 {
		a.Log.Errorf("api_repo :: GetIntraday5Quotes :: %s", err)
		return qmodels.TSMongo{}, err
	}

	body := qmodels.Intraday5TS{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&body)
	if err != nil {
		a.Log.Errorf("api_repo :: GetIntraday5Quotes :: %s", err)
		return qmodels.TSMongo{}, err
	}

	res := qmodels.TSMongo{}
	res.Set5(body)
	return res, nil
}

func (a *ApiRepo) GetIntraday15Quotes(request *http.Request) (qmodels.TSMongo, error) {
	resp, err := a.Client.Do(request)
	if err != nil || resp.StatusCode != 200 {
		a.Log.Errorf("api_repo :: GetIntraday15Quotes :: %s", err)
		return qmodels.TSMongo{}, err
	}

	body := qmodels.Intraday15TS{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&body)
	if err != nil {
		a.Log.Errorf("api_repo :: GetIntraday15Quotes :: %s", err)
		return qmodels.TSMongo{}, err
	}

	res := qmodels.TSMongo{}
	res.Set15(body)
	return res, nil
}

func (a *ApiRepo) GetIntraday30Quotes(request *http.Request) (qmodels.TSMongo, error) {
	resp, err := a.Client.Do(request)
	if err != nil || resp.StatusCode != 200 {
		a.Log.Errorf("api_repo :: GetIntraday30Quotes :: %s", err)
		return qmodels.TSMongo{}, err
	}

	body := qmodels.Intraday30TS{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&body)
	if err != nil {
		a.Log.Errorf("api_repo :: GetIntraday30Quotes :: %s", err)
		return qmodels.TSMongo{}, err
	}

	res := qmodels.TSMongo{}
	res.Set30(body)
	return res, nil
}

func (a *ApiRepo) GetIntraday60Quotes(request *http.Request) (qmodels.TSMongo, error) {
	resp, err := a.Client.Do(request)
	if err != nil || resp.StatusCode != 200 {
		a.Log.Errorf("api_repo :: GetIntraday60Quotes :: %s", err)
		return qmodels.TSMongo{}, err
	}

	body := qmodels.Intraday60TS{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&body)
	if err != nil {
		a.Log.Errorf("api_repo :: GetIntraday60Quotes :: %s", err)
		return qmodels.TSMongo{}, err
	}

	res := qmodels.TSMongo{}
	res.Set60(body)
	return res, nil
}

func (a *ApiRepo) GetDailyQuotes(request *http.Request) (qmodels.TSMongo, error) {
	resp, err := a.Client.Do(request)
	if err != nil || resp.StatusCode != 200 {
		a.Log.Errorf("api_repo :: GetDailyQuotes :: %s", err)
		return qmodels.TSMongo{}, err
	}

	body := qmodels.DailyTS{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&body)
	if err != nil {
		a.Log.Errorf("api_repo :: GetDailyQuotes :: %s", err)
		return qmodels.TSMongo{}, err
	}

	res := qmodels.TSMongo{}
	res.SetD(body)
	return res, nil
}

func (a *ApiRepo) GetWeeklyQuotes(request *http.Request) (qmodels.TSMongo, error) {
	resp, err := a.Client.Do(request)
	if err != nil || resp.StatusCode != 200 {
		a.Log.Errorf("api_repo :: GetWeeklyQuotes :: %s", err)
		return qmodels.TSMongo{}, err
	}

	body := qmodels.WeeklyTS{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&body)
	if err != nil {
		a.Log.Errorf("api_repo :: GetWeeklyQuotes :: %s", err)
		return qmodels.TSMongo{}, err
	}

	res := qmodels.TSMongo{}
	res.SetW(body)
	return res, nil
}

func (a *ApiRepo) GetMonthlyQuotes(request *http.Request) (qmodels.TSMongo, error) {
	resp, err := a.Client.Do(request)
	if err != nil || resp.StatusCode != 200 {
		a.Log.Errorf("api_repo :: GetMonthlyQuotes :: %s", err)
		return qmodels.TSMongo{}, err
	}

	body := qmodels.MonthlyTS{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&body)
	if err != nil {
		a.Log.Errorf("api_repo :: GetMonthlyQuotes :: %s", err)
		return qmodels.TSMongo{}, err
	}

	res := qmodels.TSMongo{}
	res.SetM(body)
	return res, nil
}