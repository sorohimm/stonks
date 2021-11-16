package quotes_repo

import (
	"context"
	"encoding/json"
	_ "github.com/json-iterator/go"
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

func (r *QuotesRepo) GetQuotesDB(db *mongo.Database, coll string , filter interface{}) (interface{}, error) {
	var body qmodels.TSMongo
	cursor, err := db.Collection(coll).Aggregate(context.TODO(), filter)
	if err != nil {
		r.Log.Infof("quotes repo: GetQuotesDB: %s", err)
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		if err = cursor.Decode(&body); err != nil {
			return nil, err
		}
		break
	}

	return body, nil
}

func (r *QuotesRepo) GetIntraday1Quotes(request *http.Request) (qmodels.TSMongo, error) {
	resp, err := r.Client.Do(request)
	if err != nil || resp.StatusCode != 200 {
		r.Log.Errorf("quotes repo: GetDailyQuotes request error: %s", err)
		return qmodels.TSMongo{}, err
	}

	body := qmodels.Intraday1TS{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&body)
	if err != nil {
		r.Log.Errorf("quotes repo: GetDailyQuotes decode error: %s", err)
		return qmodels.TSMongo{}, err
	}

	res := qmodels.TSMongo{}
	res.Set1(body)
	return res, nil
}

func (r *QuotesRepo) GetIntraday5Quotes(request *http.Request) (qmodels.TSMongo, error) {
	resp, err := r.Client.Do(request)
	if err != nil || resp.StatusCode != 200 {
		r.Log.Errorf("quotes repo: GetDailyQuotes request error: %s", err)
		return qmodels.TSMongo{}, err
	}

	body := qmodels.Intraday5TS{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&body)
	if err != nil {
		r.Log.Errorf("quotes repo: GetDailyQuotes decode error: %s", err)
		return qmodels.TSMongo{}, err
	}

	res := qmodels.TSMongo{}
	res.Set5(body)
	return res, nil
}

func (r *QuotesRepo) GetIntraday15Quotes(request *http.Request) (qmodels.TSMongo, error) {
	resp, err := r.Client.Do(request)
	if err != nil || resp.StatusCode != 200 {
		r.Log.Errorf("quotes repo: GetDailyQuotes request error: %s", err)
		return qmodels.TSMongo{}, err
	}

	body := qmodels.Intraday15TS{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&body)
	if err != nil {
		r.Log.Errorf("quotes repo: GetDailyQuotes decode error: %s", err)
		return qmodels.TSMongo{}, err
	}

	res := qmodels.TSMongo{}
	res.Set15(body)
	return res, nil
}

func (r *QuotesRepo) GetIntraday30Quotes(request *http.Request) (qmodels.TSMongo, error) {
	resp, err := r.Client.Do(request)
	if err != nil || resp.StatusCode != 200 {
		r.Log.Errorf("quotes repo: GetDailyQuotes request error: %s", err)
		return qmodels.TSMongo{}, err
	}

	body := qmodels.Intraday30TS{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&body)
	if err != nil {
		r.Log.Errorf("quotes repo: GetDailyQuotes decode error: %s", err)
		return qmodels.TSMongo{}, err
	}

	res := qmodels.TSMongo{}
	res.Set30(body)
	return res, nil
}

func (r *QuotesRepo) GetIntraday60Quotes(request *http.Request) (qmodels.TSMongo, error) {
	resp, err := r.Client.Do(request)
	if err != nil || resp.StatusCode != 200 {
		r.Log.Errorf("quotes repo: GetDailyQuotes request error: %s", err)
		return qmodels.TSMongo{}, err
	}

	body := qmodels.Intraday60TS{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&body)
	if err != nil {
		r.Log.Errorf("quotes repo: GetDailyQuotes decode error: %s", err)
		return qmodels.TSMongo{}, err
	}

	res := qmodels.TSMongo{}
	res.Set60(body)
	return res, nil
}

func (r *QuotesRepo) GetDailyQuotes(request *http.Request) (qmodels.TSMongo, error) {
	resp, err := r.Client.Do(request)
	if err != nil || resp.StatusCode != 200 {
		r.Log.Errorf("quotes repo: GetDailyQuotes request error: %s", err)
		return qmodels.TSMongo{}, err
	}

	body := qmodels.DailyTS{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&body)
	if err != nil {
		r.Log.Errorf("quotes repo: GetDailyQuotes decode error: %s", err)
		return qmodels.TSMongo{}, err
	}

	res := qmodels.TSMongo{}
	res.SetD(body)
	return res, nil
}

func (r *QuotesRepo) GetWeeklyQuotes(request *http.Request) (qmodels.TSMongo, error) {
	resp, err := r.Client.Do(request)
	if err != nil || resp.StatusCode != 200 {
		r.Log.Errorf("quotes repo: GetWeeklyQuotes request error: %s", err)
		return qmodels.TSMongo{}, err
	}

	body := qmodels.WeeklyTS{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&body)
	if err != nil {
		r.Log.Errorf("quotes repo: GetWeeklyQuotes decode error: %s", err)
		return qmodels.TSMongo{}, err
	}

	res := qmodels.TSMongo{}
	res.SetW(body)
	return res, nil
}

func (r *QuotesRepo) GetMonthlyQuotes(request *http.Request) (qmodels.TSMongo, error) {
	resp, err := r.Client.Do(request)
	if err != nil || resp.StatusCode != 200 {
		r.Log.Errorf("quotes repo: GetMonthlyQuotes request error: %s", err)
		return qmodels.TSMongo{}, err
	}

	body := qmodels.MonthlyTS{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&body)
	if err != nil {
		r.Log.Errorf("quotes repo: GetMonthlyQuotes decode error: %s", err)
		return qmodels.TSMongo{}, err
	}

	res := qmodels.TSMongo{}
	res.SetM(body)
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
