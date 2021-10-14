package overview_repo

import (
	"encoding/json"
	"log"
	"net/http"
	om "stonks/internal/models/overview"
)

type OverviewRepo struct {
	client http.Client
}

func (r *OverviewRepo) GetOverview(request *http.Request) (om.Overview, error) {
	resp, err := r.client.Do(request)
	if err != nil || resp.StatusCode != 200 {
		log.Print(json.NewDecoder(resp.Body))
		return om.Overview{}, err
	}

	overviewBody := om.Overview{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&overviewBody)
	if err != nil {
		log.Print(err.Error())
	}
	return overviewBody, nil
}
