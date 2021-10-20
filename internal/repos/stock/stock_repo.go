package stock_repo

import (
	"encoding/json"
	"log"
	"net/http"
	"stonks/internal/models/stock"
)


type StockRepo struct {
	client http.Client
}

var (
	structs map[string]interface{}
)

func init() {
	structs = make(map[string]interface{})
	structs["TIME_SERIES_DAILY"] = stock_models.DailyTimeSeries{}
	structs["TIME_SERIES_WEEKLY"] = stock_models.WeeklyTimeSeries{}
	structs["TIME_SERIES_MONTHLY"] = stock_models.MonthlyTimeSeries{}
}

func (r *StockRepo) GetStock(request *http.Request) (interface{}, error) {
	resp, err := r.client.Do(request)
	if err != nil || resp.StatusCode != 200 {
		log.Print(json.NewDecoder(resp.Body))
		return nil, err
	}

	seriesBody := structs[request.URL.Query().Get("function")]

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&seriesBody)
	if err != nil {
		log.Print(err.Error())
	}
	return seriesBody, nil
}
