package repos

import (
	"encoding/json"
	"log"
	"net/http"
	"stonks/internal/models/news"
)

type NewsRepo struct {
	client http.Client
}

func (r *NewsRepo) GetNews(request *http.Request) (news_models.News, error) {
	resp, err := r.client.Do(request)
	if err != nil || resp.StatusCode != 200 {
		log.Print(json.NewDecoder(resp.Body))
		return news_models.News{}, err
	}

	newsBody := news_models.News{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&newsBody)
	if err != nil {
		log.Print(err.Error())
	}
	return newsBody, nil
}
