package details_repo

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"net/http"
	"stonks/internal/models/company_details"
)

type CompanyDetailsRepo struct {
	Log    *zap.SugaredLogger
	Client *http.Client
}

var (
	models map[string]interface{}
)

func init() {
	models = make(map[string]interface{})
	models["Overview"] = &details_models.Overview{}
	models["Earnings"] = &details_models.Earnings{}
	models["IncomeStatement"] = &details_models.IncomeStatement{}
	models["BalanceSheet"] = &details_models.BalanceSheet{}
	models["CashFlow"] = &details_models.CashFlow{}
}

func (r *CompanyDetailsRepo) GetCompanyDetails(request *http.Request) (interface{}, error) {
	resp, err := r.Client.Do(request)
	if err != nil || resp.StatusCode != 200 {
		r.Log.Errorf("details repo: GetCompanyDetails request error: %s", err)
		return nil, err
	}

	body := models[request.URL.Query().Get("function")]

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&body)
	if err != nil {
		r.Log.Errorf("details_repo :: GetCompanyDetails :: %s", err)
		return nil, err
	}
	return body, nil
}

func (r *CompanyDetailsRepo) GetOverview(db *mongo.Database, filter interface{}) (interface{}, error) {
	body := details_models.OverviewMongo{}
	err := db.Collection("Overview").FindOne(context.TODO(), filter).Decode(&body)
	if err != nil {
		r.Log.Errorf("details_repo :: GetOverview :: %s", err)
		return nil, err
	}
	return body, nil
}

func (r *CompanyDetailsRepo) GetEarnings(db *mongo.Database, filter interface{}) (details_models.Earnings, error) {
	body := details_models.Earnings{}
	err := db.Collection("Earnings").FindOne(context.TODO(), filter).Decode(&body)
	if err != nil {
		r.Log.Errorf("details_repo :: GetEarnings :: %s", err)
		return details_models.Earnings{}, err
	}
	return body, nil
}

func (r *CompanyDetailsRepo) GetIncomeStatement(db *mongo.Database, filter interface{}) (details_models.IncomeStatement, error) {
	body := details_models.IncomeStatement{}
	err := db.Collection("IncomeStatement").FindOne(context.TODO(), filter).Decode(&body)
	if err != nil {
		r.Log.Errorf("details_repo :: GetIncomeStatement :: %s", err)
		return details_models.IncomeStatement{}, err
	}
	return body, nil
}

func (r *CompanyDetailsRepo) GetBalanceSheet(db *mongo.Database, filter interface{}) (details_models.BalanceSheet, error) {
	body := details_models.BalanceSheet{}
	err := db.Collection("BalanceSheet").FindOne(context.TODO(), filter).Decode(&body)
	if err != nil {
		r.Log.Errorf("details_repo :: GetBalanceSheet :: %s", err)
		return details_models.BalanceSheet{}, err
	}
	return body, nil
}

func (r *CompanyDetailsRepo) GetCashFlow(db *mongo.Database, filter interface{}) (details_models.CashFlow, error) {
	body := details_models.CashFlow{}
	err := db.Collection("CashFlow").FindOne(context.TODO(), filter).Decode(&body)
	if err != nil {
		r.Log.Errorf("details_repo :: GetCashFlow :: %s", err)
		return details_models.CashFlow{}, err
	}
	return body, nil
}

func (r *CompanyDetailsRepo) InsertCompanyDetails(collection string, db *mongo.Database, body interface{}) (interface{}, error) {
	id, err := db.Collection(collection).InsertOne(context.TODO(), body)
	if err != nil {
		r.Log.Errorf("details_repo :: InsertCompanyDetails :: %s", err)
		return nil, err
	}

	return id.InsertedID, nil
}
