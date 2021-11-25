package api_repo

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	details_models "stonks/internal/models/company_details"
	qmodels "stonks/internal/models/quotes"
)

type StocksApiRepo struct {
	Log    *zap.SugaredLogger
	Client *http.Client
}

func (r *StocksApiRepo) GetIntraday1Quotes(request *http.Request) (qmodels.TSMongo, error) {
	resp, err := r.Client.Do(request)
	if err != nil || resp.StatusCode != 200 {
		r.Log.Infof("api_repo :: GetIntraday1Quotes :: %s", err)
		return qmodels.TSMongo{}, err
	}

	body := qmodels.Intraday1TS{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&body)
	if err != nil {
		r.Log.Infof("api_repo :: GetIntraday1Quotes :: %s", err)
		return qmodels.TSMongo{}, err
	}

	res := qmodels.TSMongo{}
	res.Set1(body)

	return res, nil
}

func (r *StocksApiRepo) GetIntraday5Quotes(request *http.Request) (qmodels.TSMongo, error) {
	resp, err := r.Client.Do(request)
	if err != nil || resp.StatusCode != 200 {
		r.Log.Infof("api_repo :: GetIntraday5Quotes :: %s", err)
		return qmodels.TSMongo{}, err
	}

	body := qmodels.Intraday5TS{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&body)
	if err != nil {
		r.Log.Infof("api_repo :: GetIntraday5Quotes :: %s", err)
		return qmodels.TSMongo{}, err
	}

	res := qmodels.TSMongo{}
	res.Set5(body)

	return res, nil
}

func (r *StocksApiRepo) GetIntraday15Quotes(request *http.Request) (qmodels.TSMongo, error) {
	resp, err := r.Client.Do(request)
	if err != nil || resp.StatusCode != 200 {
		r.Log.Infof("api_repo :: GetIntraday15Quotes :: %s", err)
		return qmodels.TSMongo{}, err
	}

	body := qmodels.Intraday15TS{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&body)
	if err != nil {
		r.Log.Infof("api_repo :: GetIntraday15Quotes :: %s", err)
		return qmodels.TSMongo{}, err
	}

	res := qmodels.TSMongo{}
	res.Set15(body)

	return res, nil
}

func (r *StocksApiRepo) GetIntraday30Quotes(request *http.Request) (qmodels.TSMongo, error) {
	resp, err := r.Client.Do(request)
	if err != nil || resp.StatusCode != 200 {
		r.Log.Infof("api_repo :: GetIntraday30Quotes :: %s", err)
		return qmodels.TSMongo{}, err
	}

	body := qmodels.Intraday30TS{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&body)
	if err != nil {
		r.Log.Infof("api_repo :: GetIntraday30Quotes :: %s", err)
		return qmodels.TSMongo{}, err
	}

	res := qmodels.TSMongo{}
	res.Set30(body)

	return res, nil
}

func (r *StocksApiRepo) GetIntraday60Quotes(request *http.Request) (qmodels.TSMongo, error) {
	resp, err := r.Client.Do(request)
	if err != nil || resp.StatusCode != 200 {
		r.Log.Infof("api_repo :: GetIntraday60Quotes :: %s", err)
		return qmodels.TSMongo{}, err
	}

	body := qmodels.Intraday60TS{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&body)
	if err != nil {
		r.Log.Infof("api_repo :: GetIntraday60Quotes :: %s", err)
		return qmodels.TSMongo{}, err
	}

	res := qmodels.TSMongo{}
	res.Set60(body)

	return res, nil
}

func (r *StocksApiRepo) GetDailyQuotes(request *http.Request) (qmodels.TSMongo, error) {
	resp, err := r.Client.Do(request)
	if err != nil || resp.StatusCode != 200 {
		r.Log.Infof("api_repo :: GetDailyQuotes :: %s", err)
		return qmodels.TSMongo{}, err
	}

	body := qmodels.DailyTS{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&body)
	if err != nil {
		r.Log.Infof("api_repo :: GetDailyQuotes :: %s", err)
		return qmodels.TSMongo{}, err
	}

	res := qmodels.TSMongo{}
	res.SetD(body)

	return res, nil
}

func (r *StocksApiRepo) GetWeeklyQuotes(request *http.Request) (qmodels.TSMongo, error) {
	resp, err := r.Client.Do(request)
	if err != nil || resp.StatusCode != 200 {
		r.Log.Infof("api_repo :: GetWeeklyQuotes :: %s", err)
		return qmodels.TSMongo{}, err
	}

	body := qmodels.WeeklyTS{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&body)
	if err != nil {
		r.Log.Infof("api_repo :: GetWeeklyQuotes :: %s", err)
		return qmodels.TSMongo{}, err
	}

	res := qmodels.TSMongo{}
	res.SetW(body)

	return res, nil
}

func (r *StocksApiRepo) GetMonthlyQuotes(request *http.Request) (qmodels.TSMongo, error) {
	resp, err := r.Client.Do(request)
	if err != nil || resp.StatusCode != 200 {
		r.Log.Infof("api_repo :: GetMonthlyQuotes :: %s", err)
		return qmodels.TSMongo{}, err
	}

	body := qmodels.MonthlyTS{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&body)
	if err != nil {
		r.Log.Infof("api_repo :: GetMonthlyQuotes :: %s", err)
		return qmodels.TSMongo{}, err
	}

	res := qmodels.TSMongo{}
	res.SetM(body)

	return res, nil
}

func (r *StocksApiRepo) GetOverview(request *http.Request) (details_models.OverviewMongo, error) {
	resp, err := r.Client.Do(request)
	if err != nil || resp.StatusCode != 200 {
		r.Log.Infof("api_repo :: GetOverview :: %s", err)
		return details_models.OverviewMongo{}, err
	}

	var body details_models.Overview

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&body)
	if err != nil {
		r.Log.Infof("api_repo :: GetOverview :: %s", err)
		return details_models.OverviewMongo{}, err
	}

	var res details_models.OverviewMongo
	res.Set(body)

	return res, nil
}

func (r *StocksApiRepo) GetEarnings(request *http.Request) (details_models.EarningsMongo, error) {
	resp, err := r.Client.Do(request)
	if err != nil || resp.StatusCode != 200 {
		r.Log.Infof("api_repo :: GetEarnings :: %s", err)
		return details_models.EarningsMongo{}, err
	}

	var body details_models.Earnings

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&body)
	if err != nil {
		r.Log.Infof("api_repo :: GetEarnings :: %s", err)
		return details_models.EarningsMongo{}, err
	}

	var res details_models.EarningsMongo
	res.Set(body)

	return res, nil
}

func (r *StocksApiRepo) GetCashFlow(request *http.Request) (details_models.CashFlowMongo, error) {
	resp, err := r.Client.Do(request)
	if err != nil || resp.StatusCode != 200 {
		r.Log.Infof("api_repo :: GetCashFlow :: %s", err)
		return details_models.CashFlowMongo{}, err
	}

	var body details_models.CashFlow

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&body)
	if err != nil {
		r.Log.Infof("api_repo :: GetCashFlow :: %s", err)
		return details_models.CashFlowMongo{}, err
	}

	var res details_models.CashFlowMongo
	res.Set(body)

	r.Log.Info(body)
	r.Log.Info(res)

	return res, nil
}

func (r *StocksApiRepo) GetIncomeStatement(request *http.Request) (details_models.IncomeStatementMongo, error) {
	resp, err := r.Client.Do(request)
	if err != nil || resp.StatusCode != 200 {
		r.Log.Infof("api_repo :: GetIncomeStatement :: %s", err)
		return details_models.IncomeStatementMongo{}, err
	}

	var body details_models.IncomeStatement

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&body)
	if err != nil {
		r.Log.Infof("api_repo :: GetIncomeStatement :: %s", err)
		return details_models.IncomeStatementMongo{}, err
	}

	var res details_models.IncomeStatementMongo
	res.Set(body)
	r.Log.Info(body)
	r.Log.Info(res)
	return res, nil
}

func (r *StocksApiRepo) GetBalanceSheet(request *http.Request) (details_models.BalanceSheetMongo, error) {
	resp, err := r.Client.Do(request)
	if err != nil || resp.StatusCode != 200 {
		r.Log.Infof("api_repo :: GetIncomeStatement :: %s", err)
		return details_models.BalanceSheetMongo{}, err
	}

	var body details_models.BalanceSheet

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&body)
	if err != nil {
		r.Log.Infof("api_repo :: GetIncomeStatement :: %s", err)
		return details_models.BalanceSheetMongo{}, err
	}

	var res details_models.BalanceSheetMongo
	res.Set(body)

	r.Log.Info(body)
	r.Log.Info(res)

	return res, nil
}