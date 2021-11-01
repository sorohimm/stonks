package stock_repo

import (
	"encoding/json"
	_ "github.com/json-iterator/go"
	"log"
	"net/http"
	stock_models "stonks/internal/models/stock"
)

type StockRepo struct {
	client http.Client
}

var (
	models         map[string]interface{}
	intradayModels map[string]interface{}
)

func init() {
	models = make(map[string]interface{})
	models["TIME_SERIES_DAILY"] = stock_models.DailyTimeSeries{}
	models["TIME_SERIES_WEEKLY"] = stock_models.WeeklyTimeSeries{}
	models["TIME_SERIES_MONTHLY"] = stock_models.MonthlyTimeSeries{}

	intradayModels = make(map[string]interface{})
	intradayModels["1"] = stock_models.Intraday1TimeSeries{}
	intradayModels["5"] = stock_models.Intraday5TimeSeries{}
	intradayModels["15"] = stock_models.Intraday15TimeSeries{}
	intradayModels["30"] = stock_models.Intraday30TimeSeries{}
	intradayModels["60"] = stock_models.Intraday60TimeSeries{}
}

func GetModel(request *http.Request) interface{} {
	if request.URL.Query().Get("function") == "TIME_SERIES_INTRADAY" {
		return intradayModels[request.URL.Query().Get("interval")]
	} else {
		return models[request.URL.Query().Get("function")]
	}
}

func (r *StockRepo) GetStock(request *http.Request) (interface{}, error) {
	resp, err := r.client.Do(request)
	if err != nil || resp.StatusCode != 200 {
		log.Print(json.NewDecoder(resp.Body))
		return nil, err
	}

	seriesBody := GetModel(request)

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&seriesBody)
	if err != nil {
		log.Print(err.Error())
	}
	return seriesBody, nil
}
