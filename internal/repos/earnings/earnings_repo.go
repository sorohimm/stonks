package earnings_repo

import (
	"encoding/json"
	"log"
	"net/http"
	em "stonks/internal/models/earnings"
)

type EarningsRepo struct {
	client http.Client
}

func (r *EarningsRepo) GetEarnings(request *http.Request) (em.Earnings, error) {
	resp, err := r.client.Do(request)
	if err != nil || resp.StatusCode != 200 {
		log.Print(json.NewDecoder(resp.Body))
		return em.Earnings{}, err
	}

	earningsBody := em.Earnings{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&earningsBody)
	if err != nil {
		log.Print(err.Error())
	}
	return earningsBody, nil
}
