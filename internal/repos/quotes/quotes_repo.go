package quotes_repo

import (
	"context"
	"encoding/json"
	_ "github.com/json-iterator/go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"net/http"
	"stonks/internal/models/quotes"
	qmodels "stonks/internal/models/quotes"
)

type QuotesRepo struct {
	Log    *zap.SugaredLogger
	Client *http.Client
}

var (
	intradayModels map[string]interface{}
)

func init() {
	intradayModels = make(map[string]interface{})
	intradayModels["1min"] = quotes_models.Intraday1TS{}
	intradayModels["5min"] = quotes_models.Intraday5TS{}
	intradayModels["15min"] = quotes_models.Intraday15TS{}
	intradayModels["30min"] = quotes_models.Intraday30TS{}
	intradayModels["60min"] = quotes_models.Intraday60TS{}
}

func (r *QuotesRepo) GetIntraday(collection string, db *mongo.Database, filter interface{}) (qmodels.IntradayTSMongo, error) {
	body := qmodels.IntradayTSMongo{}
	err := db.Collection(collection).FindOne(context.TODO(), filter).Decode(&body)
	if err != nil {
		r.Log.Errorf("decode error: %s", err)
		return qmodels.IntradayTSMongo{}, err
	}
	return body, nil
}

func (r *QuotesRepo) GetDaily(ctx string, db *mongo.Database, filter interface{}) (qmodels.DailyTSMongo, error) {
	var body qmodels.DailyTSMongo
	var err error
	var res *mongo.Cursor

	if ctx == "d" {
		res, err = db.Collection("DailySeries").Aggregate(context.TODO(), filter)
		if err != nil {
			return qmodels.DailyTSMongo{}, err
		}
		err = res.Decode(&body)
	} else if ctx == "w" {
		err = db.Collection("DailySeries").FindOne(context.TODO(), filter).Decode(&body)
	}
	if err != nil {
		r.Log.Errorf("quotes repo: GetDaily error: %s", err)
		return qmodels.DailyTSMongo{}, err
	}
	return body, nil
}

func (r *QuotesRepo) GetWeekly(ctx string, db *mongo.Database, filter interface{}) (qmodels.WeeklyTSMongo, error) {
	var body qmodels.WeeklyTSMongo
	var err error
	var res *mongo.Cursor

	if ctx == "d" {
		res, err = db.Collection("WeeklySeries").Aggregate(context.TODO(), filter)
		if err != nil {
			r.Log.Info("quotes_repo: GetWeekly decode: %s", err)
			return qmodels.WeeklyTSMongo{}, err
		}
		err = res.Decode(&body)
		if err != nil {
			r.Log.Info("quotes_repo: GetWeekly decode: %s", err)
			return qmodels.WeeklyTSMongo{}, err
		}
	} else if ctx == "w" {
		err = db.Collection("WeeklySeries").FindOne(context.TODO(), filter).Decode(&body)
	}
	if err != nil {
		r.Log.Errorf("quotes repo: GetWeekly error: %s", err)
		return qmodels.WeeklyTSMongo{}, err
	}
	return body, nil
}

func (r *QuotesRepo) GetMonthly(ctx string, db *mongo.Database, filter interface{}) (qmodels.MonthlyTSMongo, error) {
	var body qmodels.MonthlyTSMongo
	var err error

	if ctx == "d" {
		var cursor *mongo.Cursor
		cursor, err = db.Collection("MonthlySeries").Aggregate(context.TODO(), filter)
		if err != nil {
			r.Log.Infof("quotes repo: GetMonthly: %s", err)
			return qmodels.MonthlyTSMongo{}, err
		}

		var results []bson.M
		if err = cursor.All(context.TODO(), &results); err != nil {
			panic(err)
		}

		if err != nil {
			r.Log.Infof("quotes repo: GetMonthly decode: %s", body)
			return qmodels.MonthlyTSMongo{}, err
		}
		return body, nil
	} else if ctx == "w" {
		err = db.Collection("MonthlySeries").FindOne(context.TODO(), filter).Decode(&body)
		if err != nil {
			r.Log.Infof("quotes repo: GetMonthly error: %s", err)
			return qmodels.MonthlyTSMongo{}, err
		}
	}

	return body, nil
}

func (r *QuotesRepo) GetIntraday1Quotes(request *http.Request) (qmodels.IntradayTSMongo, error) {
	resp, err := r.Client.Do(request)
	if err != nil || resp.StatusCode != 200 {
		r.Log.Errorf("quotes repo: GetDailyQuotes request error: %s", err)
		return qmodels.IntradayTSMongo{}, err
	}

	body := qmodels.Intraday1TS{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&body)
	if err != nil {
		r.Log.Errorf("quotes repo: GetDailyQuotes decode error: %s", err)
		return qmodels.IntradayTSMongo{}, err
	}

	res := qmodels.IntradayTSMongo{}
	res.Set1(body)
	return res, nil
}

func (r *QuotesRepo) GetIntraday5Quotes(request *http.Request) (qmodels.IntradayTSMongo, error) {
	resp, err := r.Client.Do(request)
	if err != nil || resp.StatusCode != 200 {
		r.Log.Errorf("quotes repo: GetDailyQuotes request error: %s", err)
		return qmodels.IntradayTSMongo{}, err
	}

	body := qmodels.Intraday5TS{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&body)
	if err != nil {
		r.Log.Errorf("quotes repo: GetDailyQuotes decode error: %s", err)
		return qmodels.IntradayTSMongo{}, err
	}

	res := qmodels.IntradayTSMongo{}
	res.Set5(body)
	return res, nil
}

func (r *QuotesRepo) GetIntraday15Quotes(request *http.Request) (qmodels.IntradayTSMongo, error) {
	resp, err := r.Client.Do(request)
	if err != nil || resp.StatusCode != 200 {
		r.Log.Errorf("quotes repo: GetDailyQuotes request error: %s", err)
		return qmodels.IntradayTSMongo{}, err
	}

	body := qmodels.Intraday15TS{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&body)
	if err != nil {
		r.Log.Errorf("quotes repo: GetDailyQuotes decode error: %s", err)
		return qmodels.IntradayTSMongo{}, err
	}

	res := qmodels.IntradayTSMongo{}
	res.Set15(body)
	return res, nil
}

func (r *QuotesRepo) GetIntraday30Quotes(request *http.Request) (qmodels.IntradayTSMongo, error) {
	resp, err := r.Client.Do(request)
	if err != nil || resp.StatusCode != 200 {
		r.Log.Errorf("quotes repo: GetDailyQuotes request error: %s", err)
		return qmodels.IntradayTSMongo{}, err
	}

	body := qmodels.Intraday30TS{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&body)
	if err != nil {
		r.Log.Errorf("quotes repo: GetDailyQuotes decode error: %s", err)
		return qmodels.IntradayTSMongo{}, err
	}

	res := qmodels.IntradayTSMongo{}
	res.Set30(body)
	return res, nil
}

func (r *QuotesRepo) GetIntraday60Quotes(request *http.Request) (qmodels.IntradayTSMongo, error) {
	resp, err := r.Client.Do(request)
	if err != nil || resp.StatusCode != 200 {
		r.Log.Errorf("quotes repo: GetDailyQuotes request error: %s", err)
		return qmodels.IntradayTSMongo{}, err
	}

	body := qmodels.Intraday60TS{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&body)
	if err != nil {
		r.Log.Errorf("quotes repo: GetDailyQuotes decode error: %s", err)
		return qmodels.IntradayTSMongo{}, err
	}

	res := qmodels.IntradayTSMongo{}
	res.Set60(body)
	return res, nil
}

func (r *QuotesRepo) GetDailyQuotes(request *http.Request) (qmodels.DailyTSMongo, error) {
	resp, err := r.Client.Do(request)
	if err != nil || resp.StatusCode != 200 {
		r.Log.Errorf("quotes repo: GetDailyQuotes request error: %s", err)
		return qmodels.DailyTSMongo{}, err
	}

	body := qmodels.DailyTS{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&body)
	if err != nil {
		r.Log.Errorf("quotes repo: GetDailyQuotes decode error: %s", err)
		return qmodels.DailyTSMongo{}, err
	}

	res := qmodels.DailyTSMongo{}
	res.Set(body)
	return res, nil
}

func (r *QuotesRepo) GetWeeklyQuotes(request *http.Request) (qmodels.WeeklyTSMongo, error) {
	resp, err := r.Client.Do(request)
	if err != nil || resp.StatusCode != 200 {
		r.Log.Errorf("quotes repo: GetWeeklyQuotes request error: %s", err)
		return qmodels.WeeklyTSMongo{}, err
	}

	body := qmodels.WeeklyTS{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&body)
	if err != nil {
		r.Log.Errorf("quotes repo: GetWeeklyQuotes decode error: %s", err)
		return qmodels.WeeklyTSMongo{}, err
	}

	res := qmodels.WeeklyTSMongo{}
	res.Set(body)
	return res, nil
}

func (r *QuotesRepo) GetMonthlyQuotes(request *http.Request) (qmodels.MonthlyTSMongo, error) {
	resp, err := r.Client.Do(request)
	if err != nil || resp.StatusCode != 200 {
		r.Log.Errorf("quotes repo: GetMonthlyQuotes request error: %s", err)
		return qmodels.MonthlyTSMongo{}, err
	}

	body := qmodels.MonthlyTS{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&body)
	if err != nil {
		r.Log.Errorf("quotes repo: GetMonthlyQuotes decode error: %s", err)
		return qmodels.MonthlyTSMongo{}, err
	}

	res := qmodels.MonthlyTSMongo{}
	res.Set(body)
	return res, nil
}

func (r *QuotesRepo) InsertQuotes(collection string, db *mongo.Database, body interface{}) (interface{}, error) {
	id, err := db.Collection(collection).InsertOne(context.TODO(), body)
	if err != nil {
		r.Log.Errorf("quotes repo: InsertQuotes insertion error: %s", err)
		return nil, err
	}

	return id.InsertedID, nil
}
