package repos

import (
	"encoding/json"
	"log"
	"net/http"
	"stonks/internal/models/news"
)

type NewsRepo struct {
	client *http.Client
}

func (r *NewsRepo) GetNews(req *http.Request) (news.News, error) {
	resp, err := r.client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		log.Print(json.NewDecoder(resp.Body))
		return news.News{}, err
	}

	newsBody := news.News{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&newsBody)
	if err != nil {
		log.Print(err.Error())
	}
	return newsBody, nil
}
