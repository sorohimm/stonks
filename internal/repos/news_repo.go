package repos

import (
	"encoding/json"
	"log"
	"net/http"
	"stonks/internal/config"
	"stonks/internal/models"
)
type NewsRepo struct {
}

func (r *NewsRepo) GetNewsR(cfg *config.Config, client http.Client, url string) (models.News, error) {
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	req.Header.Set("x-api-key", cfg.NKey)
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		log.Print(json.NewDecoder(resp.Body))
		return models.News{}, err
	}

	 newsBody := models.News{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&newsBody)
	if err != nil {
		log.Print(err.Error())
	}
	return newsBody, nil
}

