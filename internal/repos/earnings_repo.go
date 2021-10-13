package repos

import (
	"encoding/json"
	"log"
	"net/http"
	"stonks/internal/models/earnings"
)

type EarningsRepo struct {
	client http.Client
}

func (r *EarningsRepo) GetEarnings(request *http.Request) (earnings_models.Earnings, error) {
	resp, err := r.client.Do(request)
	if err != nil || resp.StatusCode != 200 {
		log.Print(json.NewDecoder(resp.Body))
		return earnings_models.Earnings{}, err
	}

	earningsBody := earnings_models.Earnings{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&earningsBody)
	if err != nil {
		log.Print(err.Error())
	}
	return earningsBody, nil
}
