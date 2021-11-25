package details_repo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"net/http"
	"stonks/internal/models/company_details"
)

type CompanyDetailsRepo struct {
	Log    *zap.SugaredLogger
	Client *http.Client
}

func (r *CompanyDetailsRepo) GetOverview(db *mongo.Database, filter interface{}) (details_models.OverviewMongo, error) {
	body := details_models.OverviewMongo{}

	cursor, err := db.Collection("Overview").Aggregate(context.TODO(), filter)

	ok := cursor.TryNext(context.TODO())
	if !ok {
		r.Log.Infof("details_repo :: GetOverview :: empty doc")
		return details_models.OverviewMongo{}, err
	}

	for cursor.Next(context.TODO()) {
		if err = cursor.Decode(&body); err != nil {
			r.Log.Infof("details_repo :: GetOverview :: %s", err)
			return details_models.OverviewMongo{}, err
		}
		break
	}

	if err != nil {
		r.Log.Infof("details_repo :: GetOverview :: %s", err)
		return details_models.OverviewMongo{}, err
	}
	return body, nil
}

func (r *CompanyDetailsRepo) GetEarnings(db *mongo.Database, filter interface{}) (details_models.EarningsMongo, error) {
	body := details_models.EarningsMongo{}

	cursor, err := db.Collection("Earnings").Aggregate(context.TODO(), filter)

	for cursor.Next(context.TODO()) {
		if err = cursor.Decode(&body); err != nil {
			r.Log.Infof("details_repo :: GetEarnings :: %s", err)
			return details_models.EarningsMongo{}, err
		}
		r.Log.Info(body)
		break
	}

	if err != nil {
		r.Log.Infof("details_repo :: GetEarnings :: %s", err)
		return details_models.EarningsMongo{}, err
	}
	return body, nil
}

func (r *CompanyDetailsRepo) GetIncomeStatement(db *mongo.Database, filter interface{}) (details_models.IncomeStatementMongo, error) {
	body := details_models.IncomeStatementMongo{}

	cursor, err := db.Collection("IncomeStatement").Aggregate(context.TODO(), filter)

	ok := cursor.TryNext(context.TODO())
	if !ok {
		r.Log.Infof("details_repo :: GetOverview :: empty doc")
		return details_models.IncomeStatementMongo{}, err
	}

	for cursor.Next(context.TODO()) {
		if err = cursor.Decode(&body); err != nil {
			r.Log.Infof("details_repo :: GetIncomeStatement :: %s", err)
			return details_models.IncomeStatementMongo{}, err
		}
		break
	}

	if err != nil {
		r.Log.Infof("details_repo :: GetIncomeStatement :: %s", err)
		return details_models.IncomeStatementMongo{}, err
	}

	return body, nil
}

func (r *CompanyDetailsRepo) GetBalanceSheet(db *mongo.Database, filter interface{}) (details_models.BalanceSheetMongo, error) {
	body := details_models.BalanceSheetMongo{}

	cursor, err := db.Collection("BalanceSheet").Aggregate(context.TODO(), filter)

	ok := cursor.TryNext(context.TODO())
	if !ok {
		r.Log.Infof("details_repo :: GetOverview :: empty doc")
		return details_models.BalanceSheetMongo{}, err
	}

	for cursor.Next(context.TODO()) {
		if err = cursor.Decode(&body); err != nil {
			r.Log.Infof("details_repo :: GetBalanceSheet :: %s", err)
			return details_models.BalanceSheetMongo{}, err
		}
		break
	}

	if err != nil {
		r.Log.Infof("details_repo :: GetBalanceSheet :: %s", err)
		return details_models.BalanceSheetMongo{}, err
	}
	return body, nil
}

func (r *CompanyDetailsRepo) GetCashFlow(db *mongo.Database, filter interface{}) (details_models.CashFlowMongo, error) {
	body := details_models.CashFlowMongo{}

	cursor, err := db.Collection("CashFlow").Aggregate(context.TODO(), filter)

	ok := cursor.TryNext(context.TODO())
	if !ok {
		r.Log.Infof("details_repo :: GetOverview :: empty doc")
		return details_models.CashFlowMongo{}, err
	}

	for cursor.Next(context.TODO()) {
		if err = cursor.Decode(&body); err != nil {
			r.Log.Infof("details_repo :: GetCashFlow :: %s", err)
			return details_models.CashFlowMongo{}, err
		}
		break
	}

	if err != nil {
		r.Log.Infof("details_repo :: GetCashFlow :: %s", err)
		return details_models.CashFlowMongo{}, err
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
