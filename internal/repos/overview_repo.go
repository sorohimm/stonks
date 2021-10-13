package repos

import (
	"encoding/json"
	"log"
	"net/http"
	"stonks/internal/models/overview"
)

type OverviewRepo struct {
	client http.Client
}

func (r *OverviewRepo) GetOverview(request *http.Request) (overview_models.Overview, error) {
	resp, err := r.client.Do(request)
	if err != nil || resp.StatusCode != 200 {
		log.Print(json.NewDecoder(resp.Body))
		return overview_models.Overview{}, err
	}

	overviewBody := overview_models.Overview{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&overviewBody)
	if err != nil {
		log.Print(err.Error())
	}
	return overviewBody, nil
}
