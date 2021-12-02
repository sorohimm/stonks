package details_repo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"net/http"
	"stonks/internal/config"
	"stonks/internal/models/company_details"
)

type CompanyDetailsRepo struct {
	Log    *zap.SugaredLogger
	Client *http.Client
	Config *config.Config
}

func (r *CompanyDetailsRepo) GetOverview(db *mongo.Database, filter interface{}) (details_models.OverviewMongo, error) {
	body := details_models.OverviewMongo{}

	cursor, err := db.Collection(r.Config.DetailsCollections.Overview).Aggregate(context.TODO(), filter)
	if err != nil {
		r.Log.Infof("details_repo :: GetOverview :: %s", err)
		return details_models.OverviewMongo{}, err
	}

	for cursor.Next(context.TODO()) {
		if err = cursor.Decode(&body); err != nil {
			r.Log.Infof("details_repo :: GetOverview :: %s", err)
			return details_models.OverviewMongo{}, err
		}
	}

	return body, nil
}

func (r *CompanyDetailsRepo) GetEarnings(db *mongo.Database, filter interface{}) (details_models.EarningsMongo, error) {
	body := details_models.EarningsMongo{}

	cursor, err := db.Collection(r.Config.DetailsCollections.Earnings).Aggregate(context.TODO(), filter)
	if err != nil {
		r.Log.Infof("details_repo :: GetEarnings :: %s", err)
		return details_models.EarningsMongo{}, err
	}

	for cursor.Next(context.TODO()) {
		if err = cursor.Decode(&body); err != nil {
			r.Log.Infof("details_repo :: GetEarnings :: %s", err)
			return details_models.EarningsMongo{}, err
		}
	}

	return body, nil
}

func (r *CompanyDetailsRepo) GetIncomeStatement(db *mongo.Database, filter interface{}) (details_models.IncomeStatementMongo, error) {
	body := details_models.IncomeStatementMongo{}

	cursor, err := db.Collection(r.Config.DetailsCollections.IncomeStatement).Aggregate(context.TODO(), filter)
	if err != nil {
		r.Log.Infof("details_repo :: GetIncomeStatement :: %s", err)
		return details_models.IncomeStatementMongo{}, err
	}

	for cursor.Next(context.TODO()) {
		if err = cursor.Decode(&body); err != nil {
			r.Log.Infof("details_repo :: GetIncomeStatement :: %s", err)
			return details_models.IncomeStatementMongo{}, err
		}
	}

	return body, nil
}

func (r *CompanyDetailsRepo) GetBalanceSheet(db *mongo.Database, filter interface{}) (details_models.BalanceSheetMongo, error) {
	body := details_models.BalanceSheetMongo{}

	cursor, err := db.Collection(r.Config.DetailsCollections.BalanceSheet).Aggregate(context.TODO(), filter)
	if err != nil {
		r.Log.Infof("details_repo :: GetBalanceSheet :: %s", err)
		return details_models.BalanceSheetMongo{}, err
	}

	for cursor.Next(context.TODO()) {
		if err = cursor.Decode(&body); err != nil {
			r.Log.Infof("details_repo :: GetBalanceSheet :: %s", err)
			return details_models.BalanceSheetMongo{}, err
		}
	}

	return body, nil
}

func (r *CompanyDetailsRepo) GetCashFlow(db *mongo.Database, filter interface{}) (details_models.CashFlowMongo, error) {
	body := details_models.CashFlowMongo{}

	cursor, err := db.Collection(r.Config.DetailsCollections.CashFlow).Aggregate(context.TODO(), filter)
	if err != nil {
		r.Log.Infof("details_repo :: GetCashFlow :: %s", err)
		return details_models.CashFlowMongo{}, err
	}

	for cursor.Next(context.TODO()) {
		if err = cursor.Decode(&body); err != nil {
			r.Log.Infof("details_repo :: GetCashFlow :: %s", err)
			return details_models.CashFlowMongo{}, err
		}
	}

	return body, nil
}

func (r *CompanyDetailsRepo) InsertCompanyDetails(collection string, db *mongo.Database, body interface{}) (interface{}, error) {
	id, err := db.Collection(collection).InsertOne(context.TODO(), body)
	if err != nil {
		r.Log.Infof("details_repo :: InsertCompanyDetails :: %s", err)
		return nil, err
	}

	return id.InsertedID, nil
}
