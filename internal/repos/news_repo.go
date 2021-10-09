package repos

import (
	"encoding/json"
	"log"
	"net/http"
	"stonks/internal/models"
)
type NewsRepo struct {
}

func (r *NewsRepo) GetNewsR(url string) (models.News, error) {
	client := http.Client{}
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	req.Header.Set("x-api-key", "bER6wS4I14V8eExigpXy5AC2Heinzt0A5T8_UM4DVvY")
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		log.Print(json.NewDecoder(resp.Body))
		return models.News{}, err
	}

	var newsBody models.News

	err = json.NewDecoder(resp.Body).Decode(&newsBody)
	if err != nil {
		log.Print(err.Error())
	}

	return newsBody, nil
}

